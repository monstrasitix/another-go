package view

import (
	"html/template"
	"net/http"
)

var (
	TEMPLATE = make(map[string]*template.Template)
)

func page(templ *template.Template, path string) *template.Template {
	return template.Must(template.Must(templ.Clone()).ParseFiles(path))
}

func SetupTemplates() {
	ext := template.Must(template.New("base").ParseGlob("./template/extend/*/*.go.html"))

	ext.Funcs(template.FuncMap{
		"slices": func(n int) []int {
			return make([]int, n)
		},
	})

	TEMPLATE["home"] = page(ext, "./template/home.go.html")
	TEMPLATE["builder"] = page(ext, "./template/builder.go.html")
	TEMPLATE["about"] = page(ext, "./template/about.go.html")
	TEMPLATE["pages"] = page(ext, "./template/pages.go.html")
	TEMPLATE["translations"] = page(ext, "./template/translations.go.html")
	TEMPLATE["not-found"] = page(ext, "./template/not-found.go.html")
}

func RenderTemplate(w http.ResponseWriter, name string, data any) {
	TEMPLATE[name].ExecuteTemplate(w, "base", data)
}
