package model

import "github.com/monstrasitix/finance/internal/i18n"

type ContactsModel struct {
	Page
	Sidebar
}

func NewContactsModel(path string) ContactsModel {
	return ContactsModel{
		Page: Page{
			Lang:  i18n.English.Lang,
			Title: i18n.English.TitleContacts,
		},
		Sidebar: Sidebar{
			Links: GetLinks(path, i18n.English),
		},
	}
}
