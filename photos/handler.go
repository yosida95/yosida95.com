package photos

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Handler struct {
	store StoreFactory
}

func NewHandler(store StoreFactory) *Handler {
	return &Handler{
		store: store,
	}
}

func Register(mux *http.ServeMux, h *Handler) {
	if debug {
		mux.Handle("GET /", http.FileServerFS(os.DirFS("dist/")))
	}
	mux.HandleFunc("GET /photos/{$}", h.List)
	mux.HandleFunc("GET /photos/{photo_id}", h.View)
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

type listTmplContext struct {
	Page   int64
	Last   int64
	Photos []*Photo
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

func (c *listTmplContext) URL() template.URL {
	return resolveURL(&url.URL{
		Path:     "/photos",
		RawQuery: fmt.Sprintf("page=%d", c.Page),
	})
}

func (h *Handler) List(w http.ResponseWriter, req *http.Request) {
	emitCommonHeaders(w)
	if err := req.ParseForm(); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	tx, err := h.store.Begin(ctx)
	if err != nil {
		log.Printf("Failed to start transaction: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	count, err := tx.CountPhoto()
	if err != nil {
		log.Printf("Failed to count photos: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	last := (count + 8) / 9

	var page int64
	if pageValue := req.Form.Get("page"); pageValue == "" {
		page = last
	} else {
		var err error
		page, err = strconv.ParseInt(pageValue, 10, 32)
		if err != nil || page > last {
			http.Redirect(w, req, fmt.Sprintf("/photos?page=%d", last), http.StatusFound)
			return
		}
	}

	photos, err := tx.
		FetchPhoto().
		Offset(9 * (page - 1)).
		Limit(9).
		All()
	if err != nil {
		log.Printf("Failed to fetch photos: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	if err := lookupTemplate("templates/list.tmpl").Execute(w, &listTmplContext{
		Page:   page,
		Last:   last,
		Photos: photos,
	}); err != nil {
		log.Printf("Failed to render HTML: %v", err)
	}
}

type viewTmplContext struct {
	Photo *Photo
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

func (c *viewTmplContext) URL() template.URL {
	return resolveURL(&url.URL{
		Path: path.Join("/photos", string(c.Photo.Id)),
	})
}

func (h *Handler) View(w http.ResponseWriter, req *http.Request) {
	emitCommonHeaders(w)

	name := req.PathValue("photo_id")
	var (
		id   PhotoId
		ext  string
		size string
	)
	if dot := strings.IndexByte(name, '.'); dot < 0 {
		id = PhotoId(name)
	} else {
		id, ext = PhotoId(name[:dot]), name[dot+1:]
		if dot := strings.IndexByte(ext, '.'); dot >= 0 {
			size, ext = ext[:dot], ext[dot+1:]
		}
	}

	ctx := req.Context()
	tx, err := h.store.Begin(ctx)
	if err != nil {
		log.Printf("Failed to start transaction: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	photo, err := tx.
		FetchPhoto().
		Id(id).
		First()
	if err != nil {
		if err == ErrPhotoNotFound {
			http.NotFound(w, req)
			return
		}
		log.Printf("Failed to fetch metadata: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if ext == "" {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		if err := lookupTemplate("templates/view.tmpl").Execute(w, &viewTmplContext{
			Photo: photo,
		}); err != nil {
			log.Printf("Failed to render HTML: %v", err)
		}
		return
	}
	if ext != photo.Ext() {
		http.NotFound(w, req)
		return
	}
	var key string
	switch size {
	case "":
		key = photo.Key()
	case "resized":
		key = photo.KeyResized()
	case "thumbnail":
		key = photo.KeyCropped()
	case "raw":
		http.Redirect(w, req, photo.Key(), http.StatusMovedPermanently)
		return
	default:
		http.NotFound(w, req)
		return
	}

	r, err := tx.BlobGet(key)
	if err != nil {
		if err == ErrPhotoNotFound {
			http.NotFound(w, req)
			return
		}
		log.Printf("Failed to load photo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Close()

	w.Header().Set("Cache-Control", "public, max-age=900")
	io.Copy(w, r)
}

func resolveURL(ref *url.URL) template.URL {
	base := url.URL{
		Scheme: "https",
		Host:   "yosida95.com",
	}
	return template.URL(base.ResolveReference(ref).String())
}
