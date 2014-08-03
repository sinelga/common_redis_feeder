package firstrstart

import (
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
	
	FeedRedis(*golog)
	

}
