package firstrstart

import (
	"encoding/json"
	//	"database/sql"
	"github.com/garyburd/redigo/redis"
	//	_ "github.com/mxk/go-sqlite/sqlite3"
	"domains"
	//	"fmt"
	"log/syslog"
	//	"os"
)

func FeedRedis(golog syslog.Writer, c redis.Conn) {


	var en_US_finance_menu_news = []domains.Menu{
		{Menu: "Art", OrgLink: "http://feeds.reuters.com/news/artsculture", ParsLink: "en_US:finance:news:art"},
		{Menu: "Business", OrgLink: "http://feeds.reuters.com/reuters/businessNews", ParsLink: "en_US:finance:news:business"},
		{Menu: "Business Travel", OrgLink: "http://feeds.reuters.com/ReutersBusinessTravel", ParsLink: "en_US:finance:news:businesstravel"},
		{Menu: "Company News", OrgLink: "http://feeds.reuters.com/reuters/companyNews", ParsLink: "en_US:finance:news:companynews"},
		{Menu: "Entertainment", OrgLink: "http://feeds.reuters.com/reuters/entertainment", ParsLink: "en_US:finance:news:entertainment"},
		{Menu: "Environment", OrgLink: "http://feeds.reuters.com/reuters/environment", ParsLink: "en_US:finance:news:environment"},
		{Menu: "Most Read Articles", OrgLink: "http://feeds.reuters.com/reuters/MostRead", ParsLink: "en_US:finance:news:mostread"},
	}

	var en_US_finance_menu_markets = []domains.Menu{
		{Menu: "Bankruptcy", OrgLink: "http://feeds.reuters.com/reuters/bankruptcyNews", ParsLink: "en_US:finance:markets:bankruptcy"},
		{Menu: "Bonds", OrgLink: "http://feeds.reuters.com/reuters/bondsNews", ParsLink: "en_US:finance:markets:bonds"},
	}	


	var en_US_finance_tab = []domains.Tab{
		{Tab: "News", Menus: en_US_finance_menu_news},
		{Tab: "Markets", Menus: en_US_finance_menu_markets},
	}

	var en_US_finance_localethemes = domains.LocaleThemes{Tabs: en_US_finance_tab}

	b, err := json.Marshal(en_US_finance_localethemes)
	if err != nil {

		golog.Crit(err.Error())
	}


	if _, err := c.Do("SET", "en_US:finance", b); err != nil {

		golog.Crit(err.Error())
	}

}
