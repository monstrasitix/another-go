package dashboard

import (
	"context"
	"net/http"

	"github.com/monstrasitix/finance/internal/web/model"
)

type key int

const (
	VIEW_BAG key = iota
)

func DashboardMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := model.NewDashboardModel(r.URL.Path)
		ctx := context.WithValue(r.Context(), VIEW_BAG, data)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
