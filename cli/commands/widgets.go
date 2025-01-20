package commands

import (
	"fmt"
	"os"
	"strings"
)

const (
	FILE      = "./template/extend/partial/render-widget.go.html"
	DIRECTORY = "./template/extend/widget"
)

func wrapper(content string) string {
	return fmt.Sprintf(
		`{{define "render-widget"}}<!-- -->%s%s%s{{end}}`,
		"\n",
		content,
		"\n",
	)
}

func widget(name string) string {
	return fmt.Sprintf(
		`{{if eq .Type "%s"}} {{template "widget:%s" .}} {{end}}<!-- -->`,
		name,
		name,
	)
}

func getContents(entries []os.DirEntry) []byte {
	widgets := make([]string, len(entries))

	for i, entry := range entries {
		name := entry.Name()
		widgets[i] = widget(name[:strings.IndexByte(name, '.')])
	}

	return []byte(wrapper(strings.Join(widgets, "\n")))
}

func RunWidgets(args []string) {
	entries, err := os.ReadDir(DIRECTORY)
	if err != nil {
		panic("Cannot read widget directory")
	}

	err = os.WriteFile(FILE, getContents(entries), 0644)
	if err != nil {
		panic("Cannot write widget file")
	}
}
