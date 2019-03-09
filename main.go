package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func getUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("url")

	return url
}

func scrape(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Error get" + url + "is fail")
	}

	fmt.Printf("[status] %d \n", resp.StatusCode)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("invalid status code")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("head").Each(func(i int, s *goquery.Selection) {
		s.Find("meta").Each(func(j int, m *goquery.Selection) {
			mv, _ := m.Attr("content")
			fmt.Printf("meta %s\n", mv)
		})
	})
}

func main() {
	url := getUrl()
	scrape(url)
}
