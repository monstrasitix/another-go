package model

import "github.com/monstrasitix/finance/internal/i18n"

type AboutModel struct {
	Page
	Sidebar
}

func NewAboutModel(path string) AboutModel {
	return AboutModel{
		Page: Page{
			Lang:  i18n.English.Lang,
			Title: i18n.English.TitleAbout,
		},
		Sidebar: Sidebar{
			Links: GetLinks(path, i18n.English),
		},
	}
}
