package controller

import (
	"net/http"

	"github.com/monstrasitix/finance/internal/web/model"
	"github.com/monstrasitix/finance/internal/web/view"
)

type HomeController struct {
}

func (c HomeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		view.RenderTemplate(w, "home", model.NewHomeModel(r.URL.Path))
	}
}
