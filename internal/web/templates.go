package web

import (
	"html/template"

	"github.com/monstrasitix/finance/internal/i18n"
)

var (
	TEMMPLATE = make(map[string]*template.Template)
)

func page(tmpl *template.Template, path string) *template.Template {
	return template.Must(template.Must(tmpl.Clone()).ParseFiles(path))
}

func SetupTemplates() {
	ext := template.Must(template.New("base").Funcs(template.FuncMap{
		"t": func(key string) string {
			return i18n.Text(key, key)
		},
		"slices": func(n int) []int {
			return make([]int, n)
		},
	}).ParseGlob("./template/extend/*/*.go.html"))

	TEMMPLATE["index"] = page(ext, "./template/index.go.html")
	TEMMPLATE["about"] = page(ext, "./template/about.go.html")
	TEMMPLATE["contacts"] = page(ext, "./template/contacts.go.html")
	TEMMPLATE["not-found"] = page(ext, "./template/not-found.go.html")
}
