package model

type SidebarLink struct {
	Path   string
	Label  string
	Active bool
}

func NewLink(url, path, label string) SidebarLink {
	return SidebarLink{
		Path:   path,
		Label:  label,
		Active: path == url,
	}
}

type Page struct {
	Lang         string `json:"lang"`
	Title        string `json:"title"`
	PagePath     string `json:"pagePath"`
	SidebarLinks []SidebarLink
}

func GetSidebarLinks(path string) []SidebarLink {
	return []SidebarLink{
		NewLink(path, "/", "sidebar.dashboard"),
		NewLink(path, "/about", "sidebar.about"),
	}
}
