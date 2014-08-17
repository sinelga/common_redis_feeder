package redisin

import (
	"domains"
	"encoding/json"
	"log/syslog"
	//	"fmt"
	"github.com/garyburd/redigo/redis"
	//	"log"
	//	"time"
)

func InsertIn(golog syslog.Writer, c redis.Conn, redisid string, item domains.Item) {

	bitem, err := json.Marshal(item)
	if err != nil {
		golog.Err(err.Error())

	}

	if _, err := c.Do("ZADD", redisid, item.PubDate.Unix(), bitem); err != nil {

		golog.Err(err.Error())

	}

	//	c, err := redis.Dial("tcp", ":6379")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//
	//	for _,redisidItems := range redisidItemsarr {
	//
	//	queuename := redisidItems.RedisID
	//
	//
	//	for _, item := range redisidItems.Items {
	//
	//		bitem, err := json.Marshal(item)
	//		if err != nil {
	//			fmt.Println("error:", err)
	//		}
	//		//		os.Stdout.Write(b)
	//		fmt.Println(string(bitem))
	//
	//		if pgq, err := c.Do("ZADD", queuename,item.PubDate.Unix(), bitem); err != nil {
	//			log.Fatal(err)
	//
	//		} else {
	//
	//			log.Println("in queue ", pgq)
	//
	//		}
	//
	//	}
	//}
}
