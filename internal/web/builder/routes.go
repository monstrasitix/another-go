package builder

import (
	"net/http"

	"github.com/monstrasitix/finance/internal/web/model"
	"github.com/monstrasitix/finance/internal/web/view"
)

type Widget struct {
	Type    string
	Content map[string]any
}

type Builder struct {
	Page    model.Page `json:"page"`
	Widgets []Widget   `json:"widgets"`
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	view.RenderTemplate(w, "builder", Builder{
		Page: model.Page{
			Lang:  "en",
			Title: "Builder",
		},
		Widgets: []Widget{
			{
				Type: "header",
				Content: map[string]any{
					"title":        "Build a website that preforms",
					"slogan":       "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ab, vitae. Iure tempora laborum quod deleniti. Blanditiis nostrum omnis nam qui, quaerat ad cumque ea. Similique error nemo quam illo totam.",
					"actionButton": "Get started",
				},
			},
			{
				Type: "featured",
				Content: map[string]any{
					"count": 3,
				},
			},
		},
	})
}
