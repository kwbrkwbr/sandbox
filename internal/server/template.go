package server

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init() echo.Renderer {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")), // htmlファイルのあるディレクトリ+/*.html
	}
	return t
}
