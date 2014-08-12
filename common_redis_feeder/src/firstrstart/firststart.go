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

	//	var c redis.Conn

	var en_US_finance_menu_news = []domains.Menu{
		{Menu: "Business", OrgLink: "http://feeds.reuters.com/reuters/businessNews", ParsLink: "http://feeds.reuters.com/reuters/businessNews"},
		{Menu: "Most Read Articles", OrgLink: "http://feeds.reuters.com/reuters/MostRead", ParsLink: "http://feeds.reuters.com/reuters/MostRead"},
	}

	var en_US_finance_menu_markets = []domains.Menu{
		{Menu: "Bankruptcy", OrgLink: "http://feeds.reuters.com/reuters/bankruptcyNews", ParsLink: "http://feeds.reuters.com/reuters/bankruptcyNews"},
		{Menu: "Bonds", OrgLink: "http://feeds.reuters.com/reuters/bondsNews", ParsLink: "http://feeds.reuters.com/reuters/bondsNews"},
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

	//	if c, err = redis.Dial("tcp", ":6379"); err != nil {
	//
	//		golog.Crit(err.Error())
	////		return nil
	//	} else {

	if _, err := c.Do("SET", "en_US:finance", b); err != nil {

		golog.Crit(err.Error())
	}

	//		const jsonStream = `{"locale": "en_US", "themes": "finance",[{"tab":"News",[{"menu":"Business"},{"menu":"Most Read Articles"}]]}}`

	//		db, err := sql.Open("sqlite3", "/home/juno/git/common_redis_feeder/common_redis_feeder/firststart.db")
	//		if err != nil {
	//			golog.Crit(err.Error())
	//		}
	//
	//		sqlstr := "select locale,themes,menu from menu"
	//
	//		rows, err := db.Query(sqlstr)
	//		if err != nil {
	//
	//			golog.Crit(err.Error())
	//		}
	//		defer rows.Close()
	//
	//		for rows.Next() {
	//			var locale string
	//			var themes string
	//			var menu string
	//			rows.Scan(&locale, &themes, &menu)
	//			golog.Info(locale + themes + menu)
	//
	//			redisid :=locale+":"+themes+":"+menu
	//
	//
	//		}
	//		rows.Close()

	//		return c
	//
	//	}
	//	return c
}
