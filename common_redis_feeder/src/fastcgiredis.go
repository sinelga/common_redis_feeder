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
	"strings"
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

	feeder(*golog, resp, req, callback, redisid)

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

func feeder(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, callback string, redisid string) {

	c := pool.Get()
	defer c.Close()

	if redisid != "" && callback != "" {

					var buffer bytes.Buffer
			buffer.WriteString(callback + "(")


		if strings.Index(redisid, ":news") > -1 {

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
//			var buffer bytes.Buffer
//			buffer.WriteString(callback + "(")
			buffer.Write(bout)
//			buffer.WriteString(");")
//			resp.Write(buffer.Bytes())

		} else {
			
//			golog.Info("redisid " + redisid)
			result, _ := redis.Bytes(c.Do("GET", redisid))
//			//			golog.Info(string(result))
//			var buffer bytes.Buffer
//			buffer.WriteString(callback + "(")
			buffer.Write(result)
//			buffer.WriteString(");")
//			resp.Write(buffer.Bytes())

		}
					buffer.WriteString(");")
			resp.Write(buffer.Bytes())
		
		
		//		var buffer bytes.Buffer
		//		buffer.WriteString(callback + "(")
		//		buffer.Write(result)
		//		buffer.WriteString(");")
		//
		//		resp.Write(buffer.Bytes())
	}

}
