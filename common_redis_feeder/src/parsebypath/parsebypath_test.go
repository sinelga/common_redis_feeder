package parsebypath

import (
//	"fmt"
//	"github.com/moovweb/gokogiri"
//	"io/ioutil"
	"log"
	"log/syslog"
//	"net/http"
	"testing"
//	"strings"
)

func TestParse(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}
//	client := &http.Client{}

//	xpath := []string{"//div[@class='relatedPhoto landscape']", "//div[@class='description lp_links']/p"}
	
	xpath := []string{"//div[@class='relatedPhoto landscape']", "//span[@id='articleText']/p"}
	//	link := "http://www.quotidiano.net/speciali/mondiali-2014"

	link := "http://www.reuters.com/article/2014/08/14/us-brazil-generalmotors-idUSKBN0GE1NU20140814"
	
	Parse(*golog,link,xpath)
	
//	/html/body/div[4]/div[4]/div/div[3]/div[1]/span/span[2]
	

//	req, err := http.NewRequest("GET", link, nil)
//	if err != nil {
//		golog.Err(err.Error())
//
//	}
//
//	resp, err := client.Do(req)
//	if err != nil {
//
//		golog.Err(err.Error())
//	}
//
//	defer resp.Body.Close()
//	bodybytes, err := ioutil.ReadAll(resp.Body)
//
//	if err != nil {
//		golog.Err(err.Error())
//	}
//
//	doc, _ := gokogiri.ParseHtml([]byte(bodybytes))
//	//	fmt.Println(doc)
//	res, _ := doc.Search(xpath[0])
//
//	fmt.Println(res)
//
//	for _, itemdom := range res {
//
//		if res2, err := itemdom.Search("img"); err != nil {
//
//			golog.Err(err.Error())
//		} else {
//
//			imglink := res2[0].Attr("src")
////			newitem.ImgLink = imglink
//
////fmt.Println(imglink)
//
//			if strings.HasSuffix(imglink, ".jpg") && strings.HasPrefix(link,"http://www.quotidiano.net") {	
//			
//				fmt.Println(imglink)
//			
////			imgLink := Link[0:strings.Index(newitem.Link, "-")] + "/" + newitem.ImgLink[strings.Index(newitem.ImgLink, "images"):len(newitem.ImgLink)]
//			
//			}
//
//
//
//		}
//		
//		
//		
//		
//	}

}
