package main 

import (
    "flag"
    "fmt"
    "log"
	"log/syslog"
	"github.com/garyburd/redigo/redis"
    "domains"
    "yqljsonparser"
    "parsebypath"
    "redisin"
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
	
    
  var rssLinks = []domains.RssLink {
  	
  	{Redisid: "en_US:finance:news:business",OrgLink: "select * from rss where url = 'http://feeds.reuters.com/reuters/businessNews'",LinktoParse:"https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20%22http%3A%2F%2Ffeeds.reuters.com%2Freuters%2FbusinessNews%22&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"}, 
  	{Redisid: "en_US:finance:news:mostread",OrgLink: "select * from rss where url = 'http://feeds.reuters.com/reuters/MostRead'",LinktoParse:"https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20'http%3A%2F%2Ffeeds.reuters.com%2Freuters%2FMostRead'&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"}, 
  	  	
  } 
  
  
  for _,rsslink := range rssLinks {
  	
  	
  	arrlinks :=yqljsonparser.Parser(*golog,rsslink.LinktoParse)
  	
  	for _,link := range arrlinks {
  		
  		
  		xpath := []string{"//div[@class='relatedPhoto landscape']", "//span[@id='articleText']/p"}
  		 		
  		item := parsebypath.Parse(*golog,link,xpath)
  		
  		fmt.Println(item.ImgLink)
  		fmt.Println(item.Cont)
  		
  		redisin.InsertIn(*golog, c,rsslink.Redisid,item )
  		 		
  		
  	} 	
  	  	
  }
    
    
}

