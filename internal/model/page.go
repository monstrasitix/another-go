package model

type BasePage struct {
	Title string
	Content any
}

type HomePage struct {
	Base BasePage
}

type AboutPage struct {
	Base BasePage
}