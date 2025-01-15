package model

import "github.com/monstrasitix/finance/internal/i18n"

type DashboardModel struct {
	Page    `json:"page"`
	Sidebar `json:"sidebar"`
}

func NewDashboardModel(path string) DashboardModel {
	return DashboardModel{
		Page: Page{
			Lang:  i18n.English.Lang,
			Title: i18n.English.TitleHomepage,
		},
		Sidebar: Sidebar{
			Links: GetLinks(path, i18n.English),
		},
	}
}
