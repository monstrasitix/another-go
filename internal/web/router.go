package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/monstrasitix/finance/internal/i18n"
	"github.com/monstrasitix/finance/internal/model"
)

func Router(mux *chi.Mux) {
	SetupTemplates()

	mux.Handle("/static/*",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		TEMMPLATE["index"].ExecuteTemplate(w, "base", model.Page{
			Title:        i18n.Text("title.homepage", "Homepage"),
			Lang:         "en",
			PagePath:     r.URL.Path,
			SidebarLinks: model.GetSidebarLinks(r.URL.Path),
		})
	})

	mux.Get("/contacts", func(w http.ResponseWriter, r *http.Request) {
		TEMMPLATE["contacts"].ExecuteTemplate(w, "base", model.Page{
			Title:        i18n.Text("title.contacts", "Contacts"),
			Lang:         "en",
			SidebarLinks: model.GetSidebarLinks(r.URL.Path),
		})
	})

	mux.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		TEMMPLATE["about"].ExecuteTemplate(w, "base", model.Page{
			Title:        i18n.Text("title.about", "About us"),
			Lang:         "en",
			SidebarLinks: model.GetSidebarLinks(r.URL.Path),
		})
	})

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		TEMMPLATE["not-found"].ExecuteTemplate(w, "base", nil)
	})
}
