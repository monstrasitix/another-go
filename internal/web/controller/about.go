package controller

import (
	"net/http"

	"github.com/monstrasitix/finance/internal/web/model"
	"github.com/monstrasitix/finance/internal/web/view"
)

type AboutController struct {
}

func (c AboutController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		view.RenderTemplate(w, "about", model.NewAboutModel(r.URL.Path))
	}
}
