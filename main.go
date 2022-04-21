package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Request html page
	res, err := http.Get("https://github.com/neovim/neovim/releases")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// `defer` is really important here to use
	// so it only closes the response, if
	// only there isn't any response

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s\n", res.StatusCode, res.Status)
	}

	// Load the html document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// The verbose method
	// doc.Find(".col-md-9").EachWithBreak(func(i int, s *goquery.Selection) bool {
	// 	version := strings.Split(s.Find("code").Text(), "\n")[0]
	// 	fmt.Println(title)
	// 	return false
	// })

	// The quick css selector trick to scrape content
	version := strings.Split(doc.Find("div.col-md-9 code").Text(), "\n")[0]
	fmt.Println(version)

}
