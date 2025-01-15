package dashboard

import (
	"context"
	"net/http"

	"github.com/monstrasitix/finance/internal/i18n"
	"github.com/monstrasitix/finance/internal/web/model"
	"github.com/monstrasitix/finance/internal/web/view"
)

func getViewBag(ctx context.Context, title string) model.DashboardModel {
	vb := ctx.Value(VIEW_BAG).(model.DashboardModel)
	vb.Page.Title = title

	return vb
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		view.RenderTemplate(w, "home",
			getViewBag(r.Context(), i18n.English.TitleHomepage))
	}
}

func GetAbout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		view.RenderTemplate(w, "about",
			getViewBag(r.Context(), i18n.English.TitleContacts))
	}
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		view.RenderTemplate(w, "contacts",
			getViewBag(r.Context(), i18n.English.TitleContacts))
	}
}
