package main

import (
	"firstrstart"
	"flag"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
)

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	flag.Parse()

	if *localeFlag == "" || *themesFlag == "" {

		golog.Err("tabmenufeedup check parameters")

	} else {

		if c, err := redis.Dial("tcp", ":6379"); err != nil {

			golog.Crit(err.Error())
			//		return nil
		} else {

			firstrstart.FeedRedis(*golog, c, *localeFlag, *themesFlag)

		}

	}

}
