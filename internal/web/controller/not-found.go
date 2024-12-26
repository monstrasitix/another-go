package controller

import (
	"net/http"

	"github.com/monstrasitix/finance/internal/web/view"
)

type NotFoundController struct {
}

func (c NotFoundController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	view.RenderTemplate(w, "not-found", nil)
}
