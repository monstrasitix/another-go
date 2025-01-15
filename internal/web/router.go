package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/monstrasitix/finance/internal/web/dashboard"
	"github.com/monstrasitix/finance/internal/web/view"
)

func Router(mux *chi.Mux) {
	view.SetupTemplates()

	mux.Handle("/static/*",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	mux.Route("/", func(r chi.Router) {
		r.Use(dashboard.DashboardMiddleware)

		r.HandleFunc("/", dashboard.GetIndex)
		r.HandleFunc("/about", dashboard.GetAbout)
		r.HandleFunc("/contacts", dashboard.GetContacts)
	})

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		view.RenderTemplate(w, "not-found", nil)
	})
}
