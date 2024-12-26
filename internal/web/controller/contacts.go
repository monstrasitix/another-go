package controller

import (
	"net/http"

	"github.com/monstrasitix/finance/internal/web/model"
	"github.com/monstrasitix/finance/internal/web/view"
)

type ContactsController struct {
}

func (c ContactsController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		view.RenderTemplate(w, "contacts", model.NewContactsModel(r.URL.Path))
	}
}
