package i18n

type Translation struct {
	Lang string `json:"lang"`

	// Titles
	TitleHomepage     string `json:"title.homepage"`
	TitlePages        string `json:"title.pages"`
	TitleTranslations string `json:"title.translations"`
	TitleAbout        string `json:"title.about"`

	// Sidebar
	SidebarDashboard    string `json:"sidebar.dashboard"`
	SidebarPages        string `json:"sidebar.pages"`
	SidebarAbout        string `json:"sidebar.about"`
	SidebarTranslations string `json:"sidebar.translations"`
}

var (
	English = Translation{
		Lang: "en",

		TitleHomepage:     "Homepage",
		TitlePages:        "Pages",
		TitleTranslations: "Translations",
		TitleAbout:        "About us",

		SidebarDashboard:    "Dashboard",
		SidebarPages:        "Pages",
		SidebarTranslations: "Translations",
		SidebarAbout:        "About",
	}
)
