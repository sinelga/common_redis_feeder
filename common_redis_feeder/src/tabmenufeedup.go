package main

import (
	"firstrstart"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
)

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	if c, err := redis.Dial("tcp", ":6379"); err != nil {

		golog.Crit(err.Error())
		//		return nil
	} else {
		
		
		firstrstart.FeedRedis(*golog,c)
		
		
	}

}
