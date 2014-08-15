package domains

import (
	"time"

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

type Item struct {
	PubDate time.Time
	Link	string
	Title   string
	Cont    string
	ImgLink string
}

type RssLink struct {
	
	Redisid string
	OrgLink string
	LinktoParse string
	
	
}