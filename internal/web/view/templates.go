package view

import (
	"html/template"
	"net/http"
)

var (
	TEMMPLATE = make(map[string]*template.Template)
)

func page(tmpl *template.Template, path string) *template.Template {
	return template.Must(template.Must(tmpl.Clone()).ParseFiles(path))
}

func SetupTemplates() {
	ext := template.Must(template.New("base").ParseGlob("./template/extend/*/*.go.html"))

	ext.Funcs(template.FuncMap{
		"slices": func(n int) []int {
			return make([]int, n)
		},
	})

	TEMMPLATE["home"] = page(ext, "./template/home.go.html")
	TEMMPLATE["about"] = page(ext, "./template/about.go.html")
	TEMMPLATE["contacts"] = page(ext, "./template/contacts.go.html")
	TEMMPLATE["not-found"] = page(ext, "./template/not-found.go.html")
}

func RenderTemplate(w http.ResponseWriter, name string, data any) {
	TEMMPLATE[name].ExecuteTemplate(w, "base", data)
}
