package model

import "github.com/monstrasitix/finance/internal/i18n"

type HomeModel struct {
	Page
	Sidebar
}

func NewHomeModel(path string) HomeModel {
	return HomeModel{
		Page: Page{
			Lang:  i18n.English.Lang,
			Title: i18n.English.TitleHomepage,
		},
		Sidebar: Sidebar{
			Links: GetLinks(path, i18n.English),
		},
	}
}
