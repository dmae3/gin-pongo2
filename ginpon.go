package ginpon

import (
	"net/http"
	"path"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin/render"
)

type (
	Pongo2Common struct {
		BasePath string
	}

	Pongo2 struct {
		Template *pongo2.Template
		Name     string
		Data     interface{}
	}

	Context pongo2.Context
)

var htmlContentType = []string{"text/html; charset=utf-8"}

func New(path string, isDebug bool) *Pongo2Common {
	pongo2.DefaultSet.Debug = isDebug

	return &Pongo2Common{BasePath: path}
}

func (p Pongo2Common) Instance(name string, data interface{}) render.Render {
	tpl := pongo2.Must(pongo2.FromCache(path.Join(p.BasePath, name)))

	return &Pongo2{
		Template: tpl,
		Name:     name,
		Data:     data,
	}
}

func (p Pongo2) Render(w http.ResponseWriter) error {
	p.writeContentType(w, htmlContentType)
	c := pongo2.Context(p.Data.(Context))

	return p.Template.ExecuteWriter(c, w)
}

func (p Pongo2) writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
