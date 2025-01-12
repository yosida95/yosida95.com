package photos

import (
	"embed"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

//go:embed templates/*.tmpl
var efs embed.FS

var (
	registry map[string]*template.Template
	debug    = debugT(false)
)

type debugT bool

func init() {
	registry = make(map[string]*template.Template)
	root := template.Must(buildRoot(efs))
	if err := fs.WalkDir(efs, ".", func(p string, e fs.DirEntry, err error) error {
		if err != nil || e.IsDir() {
			return err
		}
		name := e.Name()
		if !isTemplate(name) || isHelper(name) {
			return nil
		}
		registry[p] = template.Must(parseTemplate(template.Must(root.Clone()), efs, p))
		return nil
	}); err != nil {
		panic(err)
	}
}

func buildRoot(efs fs.FS) (*template.Template, error) {
	root := template.New("")
	return root, fs.WalkDir(efs, ".", func(p string, e fs.DirEntry, err error) error {
		if err != nil || e.IsDir() {
			return err
		}
		name := e.Name()
		if isTemplate(name) && isHelper(name) {
			_, err = parseTemplate(root, efs, p)
			return err
		}
		return nil
	})
}

func isTemplate(name string) bool { return path.Ext(name) == ".tmpl" }
func isHelper(name string) bool   { return strings.HasPrefix(name, "_") }

func parseTemplate(root *template.Template, efs fs.FS, name string) (*template.Template, error) {
	fh, err := efs.Open(name)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	buf := new(strings.Builder)
	if _, err := io.Copy(buf, fh); err != nil {
		return nil, err
	}
	return root.New(name).Parse(buf.String())
}

func lookupTemplate(name string) *template.Template {
	if debug {
		efs := os.DirFS("./photos")
		root := template.Must(buildRoot(efs))
		return template.Must(parseTemplate(root, efs, name))
	}
	return registry[name]
}
