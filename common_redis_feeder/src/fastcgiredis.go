package main

import (
	"bytes"
	"domains"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/fcgi"
	"time"
	//	"bytes"
	"encoding/json"
	//	"fmt"
//	"strings"
)

type FastCGIServer struct{}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

var (
	pool *redis.Pool
)

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	callback := req.Header.Get("X-CALLBACK")
	redisid := req.Header.Get("X-REDISID")
	redisdo := req.Header.Get("X-REDISDO")

	feeder(*golog, resp, req, callback, redisid,redisdo)

}

func main() {

	pool = newPool()

	listener, err := net.Listen("tcp", "127.0.0.1:8010")
	if err != nil {
		log.Fatal(err)
	}
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}

func feeder(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, callback string, redisid string,redisdo string) {

	c := pool.Get()
	defer c.Close()

	if redisid != "" && callback != "" && redisdo !="" {

		var buffer bytes.Buffer
		buffer.WriteString(callback + "(")

		if redisdo == "ZREVRANGE" {

			result, _ := redis.Strings(c.Do("ZREVRANGE", redisid, "0", "25"))

			var itemsarr []domains.Item

			for _, item := range result {

				var itemobj domains.Item

				b := []byte(item)
				err := json.Unmarshal(b, &itemobj)
				if err != nil {
					log.Fatal(err)
				}
				itemsarr = append(itemsarr, itemobj)

			}

			bout, err := json.Marshal(itemsarr)
			if err != nil {
				log.Fatal(err)
			}

			buffer.Write(bout)

		} else {

			result, _ := redis.Bytes(c.Do(redisdo, redisid))
			buffer.Write(result)

		}
		buffer.WriteString(");")
		resp.Write(buffer.Bytes())

	}

}
