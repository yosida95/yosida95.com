package handler

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/yosida95/yosida95.com/photos/pkg/photos"
	"github.com/yosida95/yosida95.com/photos/pkg/store"
)

type Handler struct {
	store store.StoreFactory
}

func NewHandler(store store.StoreFactory) *Handler {
	return &Handler{
		store: store,
	}
}

func Register(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/photos", h.List)
	mux.HandleFunc("/photos/", h.View)
	mux.HandleFunc("/photos/upload.xml", h.Upload)
}

var commonHeaders = map[string]string{
	"X-Content-Type-Options": "nosniff",
	"X-Frame-Options":        "SAMEORIGIN",
	"X-XSS-Protection":       "1; mode=block",
}

func emitCommonHeaders(w http.ResponseWriter) {
	h := w.Header()
	for k, v := range commonHeaders {
		h.Set(k, v)
	}
}

var listTmpl = template.Must(
	template.New("list").Parse(string(MustAsset("templates/list.tmpl"))))

type listTmplContext struct {
	Page   int64
	Last   int64
	Photos []*photos.Photo
}

func (c *listTmplContext) PrevPage() int64 {
	return c.Page - 1
}

func (c *listTmplContext) NextPage() int64 {
	if c.Page < c.Last {
		return c.Page + 1
	}
	return 0
}

func (h *Handler) List(w http.ResponseWriter, req *http.Request) {
	emitCommonHeaders(w)
	if req.Method != http.MethodGet {
		w.Header().Set("Methods", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := req.ParseForm(); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	store, err := h.store.Begin(ctx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer store.Rollback()

	count, err := store.PhotoCount()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	last := (count + 8) / 9

	page, err := strconv.ParseInt(req.Form.Get("page"), 10, 32)
	if err != nil || page > last {
		http.Redirect(w, req, fmt.Sprintf("/photos?page=%d", last), http.StatusFound)
		return
	}

	photos, err := store.
		PhotoFetch().
		Offset(9 * (page - 1)).
		Limit(9).
		All()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if len(photos) == 0 {
		http.NotFound(w, req)
		return
	}

	w.WriteHeader(http.StatusOK)
	listTmpl.Execute(w, &listTmplContext{
		Page:   page,
		Last:   last,
		Photos: photos,
	})
}

var viewTmpl = template.Must(
	template.New("view").Parse(string(MustAsset("templates/view.tmpl"))))

type viewTmplContext struct {
	Photo *photos.Photo
}

func (c *viewTmplContext) Description() string {
	if c.Photo.Comment != "" {
		if max := 50; utf8.RuneCountInString(c.Photo.Comment) > max {
			var i int
			for r := 0; r < max; r++ {
				_, size := utf8.DecodeRuneInString(c.Photo.Comment[i:])
				i += size
			}
			return c.Photo.Comment[:i] + "..."
		}
		return c.Photo.Comment
	}
	return c.Photo.Id.String()
}

func (h *Handler) View(w http.ResponseWriter, req *http.Request) {
	emitCommonHeaders(w)
	if req.Method != http.MethodGet {
		w.Header().Set("Methods", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	base, name := path.Split(req.URL.Path)
	if base != "/photos/" {
		http.NotFound(w, req)
		return
	}
	if name == "" {
		http.Redirect(w, req, "/photos", http.StatusMovedPermanently)
		return
	}

	var id photos.PhotoId
	var ext, size string
	if dot := strings.IndexByte(name, '.'); dot < 0 {
		id = photos.PhotoId(name)
	} else {
		id, ext = photos.PhotoId(name[:dot]), name[dot+1:]
		if dot := strings.IndexByte(ext, '.'); dot >= 0 {
			size, ext = ext[:dot], ext[dot+1:]
		}
	}

	ctx := req.Context()
	store, err := h.store.Begin(ctx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer store.Rollback()

	photo, err := store.
		PhotoFetch().
		Id(id).
		First()
	if err != nil {
		if err == photos.ErrPhotoNotFound {
			http.NotFound(w, req)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if ext == "" {
		w.WriteHeader(http.StatusOK)
		viewTmpl.Execute(w, &viewTmplContext{
			Photo: photo,
		})
		return
	}
	if ext != photo.Ext() {
		http.NotFound(w, req)
		return
	}
	var key string
	switch size {
	default:
		http.NotFound(w, req)
		return
	case "raw":
		http.Redirect(w, req, "/photos/"+photo.Key(), http.StatusMovedPermanently)
		return
	case "":
		key = photo.Key()
	case "resized":
		key = photo.KeyResized()
	case "thumbnail":
		key = photo.KeyCropped()
	}

	r, err := store.BlobGet(key)
	if err != nil {
		if err == photos.ErrPhotoNotFound {
			http.NotFound(w, req)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Close()

	w.WriteHeader(http.StatusOK)
	io.Copy(w, r)
}

func (h *Handler) Upload(w http.ResponseWriter, req *http.Request) {
}
