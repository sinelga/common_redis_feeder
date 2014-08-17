package parsebypath

import (
	"domains"
	"github.com/moovweb/gokogiri"
	"io/ioutil"
	"log/syslog"
	"net/http"
	//	"strings"
//	"fmt"
)

func Parse(golog syslog.Writer, item domains.Item, xpath []string) domains.Item {

	client := &http.Client{}

	req, err := http.NewRequest("GET", item.Link, nil)
	if err != nil {
		golog.Err(err.Error())

	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.0; WOW64; rv:24.0) Gecko/20100101 Firefox/24.0")
	resp, err := client.Do(req)
	if err != nil {

		golog.Err(err.Error())
	}

	defer resp.Body.Close()
	bodybytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		golog.Err(err.Error())
	}

	doc, _ := gokogiri.ParseHtml([]byte(bodybytes))

	res, _ := doc.Search(xpath[0])

	for _, itemdom := range res {

		if res2, err := itemdom.Search("img"); err != nil {

			golog.Err(err.Error())
		} else {
			imglink := res2[0].Attr("src")

//			fmt.Println(imglink)
			item.ImgLink = imglink

		}

	}

	res, _ = doc.Search(xpath[1])

	var content string

	for i, itemdom := range res {

//		fmt.Println(itemdom.Content())

		if i == 0 {

			content = itemdom.Content()

		} else {

			content = content + " " + itemdom.Content()

		}

		item.Cont = content 

	}
	
	return item

	//	var redisidItemsarrret []domains.RedisidItems
	//
	//	for _, redisidItems := range redisidItemsarr {
	//
	//		var redisidItemsret domains.RedisidItems
	//
	//		redisidItemsret.RedisID = redisidItems.RedisID
	//
	//		var returnitemsarr []domains.Item
	//
	//		for _, item := range redisidItems.Items {
	//
	//			newitem := item
	//
	//			req, err := http.NewRequest("GET", item.Link, nil)
	//			if err != nil {
	//				golog.Err(err.Error())
	//
	//			}
	//
	//			resp, err := client.Do(req)
	//			if err != nil {
	//
	//				golog.Err(err.Error())
	//			}
	//
	//			defer resp.Body.Close()
	//			bodybytes, err := ioutil.ReadAll(resp.Body)
	//
	//			if err != nil {
	//				golog.Err(err.Error())
	//			}
	//			doc, _ := gokogiri.ParseHtml([]byte(bodybytes))
	//			res, _ := doc.Search(xpath[0])
	//
	//			for _, itemdom := range res {
	//
	//				if res2, err := itemdom.Search("img"); err != nil {
	//
	//					golog.Err(err.Error())
	//				} else {
	//
	//					imglink := res2[0].Attr("src")
	//					newitem.ImgLink = imglink
	//
	//					if strings.HasSuffix(newitem.ImgLink, ".jpg") && strings.HasPrefix(item.Link, "http://www.quotidiano.net") {
	//
	//						//			if strings.HasSuffix(newitem.ImgLink, ".jpg") {
	//						//				imgLink := newitem.Link[0:strings.Index(newitem.Link, "-")] + "/" + newitem.ImgLink[strings.Index(newitem.ImgLink, "images"):len(newitem.ImgLink)]
	//						//				newitem.ImgLink = imgLink
	//						newitem.ImgLink = "http://www.quotidiano.net" + newitem.ImgLink
	//						//						returnitemsarr = append(returnitemsarr, newitem)
	//
	//						res, _ = doc.Search(xpath[1])
	//
	//						var cont string
	//						for i, itemdom := range res {
	//
	//							if i == 0 {
	//								cont = itemdom.Content()
	//							} else {
	//
	//								cont = cont + " " + itemdom.Content()
	//							}
	//							newitem.Cont = cont
	//
	//						}
	//
	//						if newitem.ImgLink != "" && newitem.Cont != "" {
	//							returnitemsarr = append(returnitemsarr, newitem)
	//
	//						}
	//
	//					}
	//
	//				}
	//			}
	//
	//			redisidItemsret.Items = returnitemsarr
	//		}
	//
	//		redisidItemsarrret = append(redisidItemsarrret, redisidItemsret)
	//
	//	}
	//	return redisidItemsarrret
}
