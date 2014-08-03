package domains

import (

)

type Menu struct {
	Menu, OrgLink, ParsLink string
}

type Tab struct {
	Tab   string
	Menus []Menu
}

type LocaleThemes struct {
	Tabs []Tab
}