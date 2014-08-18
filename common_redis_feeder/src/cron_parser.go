package main

import (
	"domains"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
	"parsebypath"
	"redisin"
	"yqljsonparser"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)

	}

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		//		log.Fatal(err)
		golog.Err(err.Error())
	}

	var rssLinks = []domains.RssLink{
		
		{Redisid: "en_US:finance:news:art", OrgLink: "select * from rss where url = 'http://feeds.reuters.com/news/artsculture'", LinktoParse: "https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20'http%3A%2F%2Ffeeds.reuters.com%2Fnews%2Fartsculture'&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"},
		{Redisid: "en_US:finance:news:business", OrgLink: "select * from rss where url = 'http://feeds.reuters.com/reuters/businessNews'", LinktoParse: "https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20%22http%3A%2F%2Ffeeds.reuters.com%2Freuters%2FbusinessNews%22&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"},
		{Redisid: "en_US:finance:news:mostread", OrgLink: "select * from rss where url = 'http://feeds.reuters.com/reuters/MostRead'", LinktoParse: "https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20'http%3A%2F%2Ffeeds.reuters.com%2Freuters%2FMostRead'&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"},
		{Redisid: "en_US:finance:markets:bankruptcy", OrgLink: "select * from rss where url = 'http://feeds.reuters.com/reuters/bankruptcyNews'", LinktoParse: "https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20'http%3A%2F%2Ffeeds.reuters.com%2Freuters%2FbankruptcyNews'&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"},
		{Redisid: "en_US:finance:markets:bonds", OrgLink: "select * from rss where url = 'http://feeds.reuters.com/reuters/bondsNews'", LinktoParse: "https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20'http%3A%2F%2Ffeeds.reuters.com%2Freuters%2FbondsNews'&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"},
	
	}

	for _, rsslink := range rssLinks {

		arrlinks := yqljsonparser.Parser(*golog, rsslink.LinktoParse)

		for _, link := range arrlinks {

			xpath := []string{"//div[@class='relatedPhoto landscape']", "//span[@id='articleText']/p"}

			item := parsebypath.Parse(*golog, link, xpath)

			fmt.Println(item.ImgLink)
			fmt.Println(item.Cont)

			if item.ImgLink != "" {

				redisin.InsertIn(*golog, c, rsslink.Redisid, item)

			}

		}

	}

}
