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
			http.FileServer(http.Dir("./public"))))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		TEMMPLATE["index"].ExecuteTemplate(w, "base", model.Page{
			Title:        i18n.Text("title.homepage", "Homepage"),
			Lang:         "en",
			Path:         r.URL.Path,
			SidebarLinks: model.GetSidebarLinks(),
		})
	})

	mux.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		TEMMPLATE["about"].ExecuteTemplate(w, "base", model.Page{
			Title:        i18n.Text("title.about", "About us"),
			Lang:         "en",
			Path:         r.URL.Path,
			SidebarLinks: model.GetSidebarLinks(),
		})
	})
}
