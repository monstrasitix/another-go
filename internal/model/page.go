package model

type SidebarLink struct {
	Path  string
	Label string
}

type Page struct {
	Lang         string `json:"lang"`
	Title        string `json:"title"`
	Path         string `json:"path"`
	SidebarLinks []SidebarLink
}

func GetSidebarLinks() []SidebarLink {
	return []SidebarLink{
		{"/", "sidebar.home"},
		{"/about", "sidebar.about"},
	}
}
