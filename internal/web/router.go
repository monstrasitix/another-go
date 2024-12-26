package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/monstrasitix/finance/internal/web/controller"
	"github.com/monstrasitix/finance/internal/web/view"
)

func Router(mux *chi.Mux) {
	view.SetupTemplates()

	mux.Handle("/static/*",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	mux.Handle("/", controller.HomeController{})
	mux.Handle("/about", controller.AboutController{})
	mux.Handle("/contacts", controller.ContactsController{})

	mux.NotFound(controller.NotFoundController{}.ServeHTTP)
}
