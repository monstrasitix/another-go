package model

import "github.com/monstrasitix/finance/internal/i18n"

type Sidebar struct {
	Links []Link `json:"links"`
}

type Link struct {
	Path     string `json:"path"`
	Label    string `json:"label"`
	Icon     string `json:"icon"`
	IsActive bool   `json:"isActive"`
}

func GetLinks(path string, t i18n.Translation) []Link {
	links := []Link{
		{
			Path:  "/",
			Label: t.SidebarDashboard,
			Icon:  "bx bxs-dashboard",
		},
		{
			Path:  "/contacts",
			Label: t.SidebarContacts,
			Icon:  "bx bxs-contact",
		},
		{
			Path:  "/about",
			Label: t.SidebarAbout,
			Icon:  "bx bx-info-circle",
		},
	}

	for i, link := range links {
		links[i].IsActive = link.Path == path
	}

	return links
}
