package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/monstrasitix/finance/internal/model"
)

func Router(mux *chi.Mux) {
	SetupTemplates()

	mux.Handle("/static/*",
		http.StripPrefix("/static/",
		http.FileServer(http.Dir("./public"))))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		TEMMPLATE["index"].ExecuteTemplate(w, "base", model.HomePage{
			Base: model.BasePage{
				Title: "Homepage",
			},
		})
	})
	
	mux.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		TEMMPLATE["about"].ExecuteTemplate(w, "base", model.AboutPage{
			Base: model.BasePage{
				Title: "About",
			},
		})
	})
}
