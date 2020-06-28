package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

//echo.Renderer interface
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()

	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo World!!")
	})

	e.GET("/index", func(c echo.Context) error {
		// テンプレートに渡す値

		data := struct {
			Content_a string
		}{

			Content_a: "雨が降っています。",
		}
		return c.Render(http.StatusOK, "index", data)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
