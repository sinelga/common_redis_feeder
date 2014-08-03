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
        MaxIdle: 3,
        IdleTimeout: 240 * time.Second,
        Dial: func () (redis.Conn, error) {
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

	locale := req.Header.Get("X-LOCALE")
	themes := req.Header.Get("X-THEMES")
	//	pathinfo := req.Header.Get("X-PATHINFO")
//	pathinfoUrl := req.URL

	feeder(*golog, resp, req, locale, themes)

}

func main() {

	log.Println("Server Start")
	
	pool = newPool()

	listener, err := net.Listen("tcp", "127.0.0.1:8010")
	if err != nil {
		log.Fatal(err)
	}
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}

func feeder(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, locale string, themes string) {

	c := pool.Get()
    defer c.Close()	


	redisid := locale + ":" + themes

	if redisid == "en_US:finance" {

		golog.Info("redisid " + redisid)
		//		c.Do("GET", redisid )
		webstruct, _ := redis.Bytes(c.Do("GET", redisid))
		golog.Info(string(webstruct))

	}

	//	webstruct, _ := redis.Bytes(c.Do("GET", redisid ))
	//	c.Do("GET", redisid )
	//
	//	golog.Info(string(webstruct))

	//	redisidconv := strings.Replace(redisid, "%3A", ":", -1)
	//
	//	golog.Info("redisid-> " + redisidconv)
	//	c, err := redis.Dial("tcp", ":6379")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	if redisid == "" {
	//
	//		redisid = "it_IT:news:Home"
	//
	//	}
	//
	//	bitem, _ := redis.Strings(c.Do("ZREVRANGE", redisidconv, "0", "25"))
	//
	//	var itemsarr []domains.Item
	//
	//	for _, item := range bitem {
	//
	//		var itemobj domains.Item
	//
	//		b := []byte(item)
	//		err := json.Unmarshal(b, &itemobj)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		itemsarr = append(itemsarr, itemobj)
	//
	//	}
	//
	//	bout, err := json.Marshal(itemsarr)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	var buffer bytes.Buffer
	//	buffer.WriteString(callback + "(")
	//	buffer.Write(bout)
	//	buffer.WriteString(");")
	//
	//	resp.Write(buffer.Bytes())
}
