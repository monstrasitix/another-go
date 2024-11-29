package model

type SidebarLink struct {
	Path   string
	Label  string
	Icon   string
	Active bool
}

type Page struct {
	Lang         string `json:"lang"`
	Title        string `json:"title"`
	PagePath     string `json:"pagePath"`
	SidebarLinks []SidebarLink
}

func GetSidebarLinks(path string) []SidebarLink {
	links := []SidebarLink{
		{
			Path:  "/",
			Label: "sidebar.dashboard",
			Icon:  "bx bxs-dashboard",
		},
		{
			Path:  "/contacts",
			Label: "sidebar.contacts",
			Icon:  "bx bxs-contact",
		},
		{
			Path:  "/about",
			Label: "sidebar.about",
			Icon:  "bx bx-info-circle",
		},
	}

	for i, link := range links {
		links[i].Active = link.Path == path
	}

	return links
}
