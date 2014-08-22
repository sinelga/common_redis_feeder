package firstrstart

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"domains"
//	"fmt"
	"log/syslog"
	//	"os"
	"database/sql"
	_ "github.com/mxk/go-sqlite/sqlite3"
)

func FeedRedis(golog syslog.Writer, c redis.Conn, locale string, themes string) {

	db, err := sql.Open("sqlite3", "/home/juno/git/common_redis_feeder/common_redis_feeder/firststart.db")
	if err != nil {

		golog.Crit(err.Error())
	}

	sqlstr := "select Tab,Menu,Redisid from feedlinks where Locale='" + locale + "' and Themes='" + themes + "'"

	rows, err := db.Query(sqlstr)
	if err != nil {

		golog.Crit(err.Error())
	}
	defer rows.Close()

	tabsroot := domains.Tabs{}

	tabs := []domains.Tab{}

	for rows.Next() {

		tab := domains.Tab{}
		menu := domains.Menu{}

		rows.Scan(&tab.Tab, &menu.Menu, &menu.Redisid)

		if len(tabs) == 0 || tabs[len(tabs)-1].Tab != tab.Tab {

			tab.Menus = append(tab.Menus, menu)
			tabs = append(tabs, tab)

		} else {

			tabs[len(tabs)-1].Menus = append(tabs[len(tabs)-1].Menus, menu)
		}

	}
	rows.Close()

	tabsroot.Tabs = tabs

	buf, err := json.Marshal(tabsroot)
	if err != nil {
		//		log.Fatal(err)
		golog.Crit(err.Error())
	}
//	fmt.Println(string(buf))

	//	var en_US_finance_menu_news = []domains.Menu{
	//		{Menu: "Art", OrgLink: "http://feeds.reuters.com/news/artsculture", ParsLink: "en_US:finance:news:art"},
	//		{Menu: "Business", OrgLink: "http://feeds.reuters.com/reuters/businessNews", ParsLink: "en_US:finance:news:business"},
	//		{Menu: "Business Travel", OrgLink: "http://feeds.reuters.com/ReutersBusinessTravel", ParsLink: "en_US:finance:news:businesstravel"},
	//		{Menu: "Company News", OrgLink: "http://feeds.reuters.com/reuters/companyNews", ParsLink: "en_US:finance:news:companynews"},
	//		{Menu: "Entertainment", OrgLink: "http://feeds.reuters.com/reuters/entertainment", ParsLink: "en_US:finance:news:entertainment"},
	//		{Menu: "Environment", OrgLink: "http://feeds.reuters.com/reuters/environment", ParsLink: "en_US:finance:news:environment"},
	//		{Menu: "Most Read Articles", OrgLink: "http://feeds.reuters.com/reuters/MostRead", ParsLink: "en_US:finance:news:mostread"},
	//	}
	//
	//	var en_US_finance_menu_markets = []domains.Menu{
	//		{Menu: "Bankruptcy", OrgLink: "http://feeds.reuters.com/reuters/bankruptcyNews", ParsLink: "en_US:finance:markets:bankruptcy"},
	//		{Menu: "Bonds", OrgLink: "http://feeds.reuters.com/reuters/bondsNews", ParsLink: "en_US:finance:markets:bonds"},
	//	}
	//
	//
	//	var en_US_finance_tab = []domains.Tab{
	//		{Tab: "News", Menus: en_US_finance_menu_news},
	//		{Tab: "Markets", Menus: en_US_finance_menu_markets},
	//	}
	//
	//	var en_US_finance_localethemes = domains.LocaleThemes{Tabs: en_US_finance_tab}
	//
	//		b, err := json.Marshal(en_US_finance_localethemes)
	//		if err != nil {
	//
	//			golog.Crit(err.Error())
	//		}

	if _, err := c.Do("SET", "en_US:finance", buf); err != nil {

		golog.Crit(err.Error())
	}

}
