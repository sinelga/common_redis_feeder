package firstrstart

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
	"testing"
)

func TestFeedRedis(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	if c, err := redis.Dial("tcp", ":6379"); err != nil {

		golog.Crit(err.Error())
		//		return nil
	} else {

		FeedRedis(*golog, c, "en_US", "finance")

	}

}
