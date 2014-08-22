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


	if _, err := c.Do("SET", "en_US:finance", buf); err != nil {

		golog.Crit(err.Error())
	}

}
