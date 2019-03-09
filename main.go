package main

import (
	"encoding/csv"
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

	file := resultCsvFile()
	writer := csv.NewWriter(file)
	doc.Find("head").Each(func(i int, s *goquery.Selection) {
		s.Find("meta").Each(func(j int, m *goquery.Selection) {
			mv, _ := m.Attr("content")
			fmt.Println(mv)
		})
	})
}

func resultCsvFile() *os.File {
	file, err := os.OpenFile("tmp/result.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	err = file.Truncate(0)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func main() {
	url := getUrl()
	scrape(url)
}
