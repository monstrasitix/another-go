package web

import (
	"html/template"
)

var (
	TEMMPLATE = make(map[string]*template.Template)
)

func page(tmpl *template.Template, path string) *template.Template {
	return template.Must(template.Must(tmpl.Clone()).ParseFiles(path))
}

func SetupTemplates() {
	ext := template.Must(template.New("base").ParseGlob("./template/extend/*/*.go.html"))

	TEMMPLATE["index"] = page(ext, "./template/index.go.html")
	TEMMPLATE["about"] = page(ext, "./template/about.go.html")
}
