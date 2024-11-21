package i18n

var (
	texts = map[string]string{
		"title.homepage": "Homepage",
		"title.about":    "About us",
	}
)

func Text(key, fallback string) string {
	if value, exists := texts[key]; exists {
		return value
	}

	return fallback
}
