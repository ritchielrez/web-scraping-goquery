package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	// "strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Request html page
	res, err := http.Get("https://github.com/neovim/neovim/releases")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s\n", res.StatusCode, res.Status)
	}

	// Load the html document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	doc.Find(".col-md-9").EachWithBreak(func(i int, s *goquery.Selection) bool {
		title := strings.Split(s.Find("code").Text(), "\n")[0]
		fmt.Println(title)
		return false
	})
}
