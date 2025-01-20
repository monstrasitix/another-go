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
			getViewBag(r.Context(), i18n.English.TitleAbout))
	}
}

func GetPages(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		view.RenderTemplate(w, "pages",
			getViewBag(r.Context(), i18n.English.TitlePages))
	}
}

func GetTranslations(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		view.RenderTemplate(w, "translations",
			getViewBag(r.Context(), i18n.English.TitleTranslations))
	}
}
