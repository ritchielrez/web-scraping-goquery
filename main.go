package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Request html page
	res, err := http.Get("https://metalsucks.net")
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

    // Find the review items
  doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the title
    title := s.Find("a").Text()
    fmt.Printf("Review %d: %s\n", i, title)
  })
}
