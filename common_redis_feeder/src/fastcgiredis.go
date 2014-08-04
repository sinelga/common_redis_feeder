package main

import (
	//	"bytes"
	//	"domains"
	//	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/fcgi"
	//	"net/url"
	"bytes"
	"time"
	//	"errors"
	//	"sync"
	//	"firstrstart"
	//	"strings"
)

//var startOnce sync.Once
//var c redis.Conn
//var rederr error

type FastCGIServer struct{}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "104.131.209.134:6379")
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

	//	locale := req.Header.Get("X-LOCALE")
	//	themes := req.Header.Get("X-THEMES")

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


	golog.Info("redisid " + redisid)
	result, _ := redis.Bytes(c.Do("GET", redisid))
	golog.Info(string(result))

	var buffer bytes.Buffer
	buffer.WriteString(callback + "(")
	buffer.Write(result)
	buffer.WriteString(");")

	resp.Write(buffer.Bytes())


}
