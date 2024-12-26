package i18n

type Translation struct {
	Lang string `json:"lang"`

	// Titles
	TitleHomepage string `json:"title.homepage"`
	TitleContacts string `json:"title.contacts"`
	TitleAbout    string `json:"title.about"`

	// Sidebar
	SidebarDashboard string `json:"sidebar.dashboard"`
	SidebarContacts  string `json:"sidebar.contacts"`
	SidebarAbout     string `json:"sidebar.about"`
}

var (
	English = Translation{
		Lang: "en",

		TitleHomepage: "Homepage",
		TitleContacts: "Cnontacts",
		TitleAbout:    "About us",

		SidebarDashboard: "Dashboard",
		SidebarContacts:  "Contacts",
		SidebarAbout:     "About",
	}
)
