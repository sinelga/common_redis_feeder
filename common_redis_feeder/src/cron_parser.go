package main 

import (
    "flag"
    "fmt"
    "log"
	"log/syslog"
    "domains"
    "yqljsonparser"
    "parsebypath"
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
    
  var rssLinks = []domains.RssLink {
  	
  	{Redisid: "en_US:finance:news:business",OrgLink: "select * from rss where url = 'http://feeds.reuters.com/reuters/businessNews'",LinktoParse:"https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20rss%20where%20url%20%3D%20%22http%3A%2F%2Ffeeds.reuters.com%2Freuters%2FbusinessNews%22&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"}, 
  	  	
  } 
  
  
  for _,rsslink := range rssLinks {
  	
  	
  	arrlinks :=yqljsonparser.Parser(*golog,rsslink.LinktoParse)
  	
  	for _,link := range arrlinks {
  		
  		
  		fmt.Println(link)
  		
  		xpath := []string{"//div[@class='relatedPhoto landscape']", "//span[@id='articleText']/p"}
  		
  		
  		parsebypath.Parse(*golog,link,xpath) 		
  		
  	} 	
  	
  	
  }
    
    
}
