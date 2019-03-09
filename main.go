package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	file, err := os.Create("tmp/result.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{"url", "property", "name", "content", "timestamp"}
	writer.Write(headers)
	doc.Find("head").Each(func(i int, s *goquery.Selection) {
		s.Find("meta").Each(func(j int, m *goquery.Selection) {
			property, _ := m.Attr("property")
			name, _ := m.Attr("name")
			content, _ := m.Attr("content")
			information := []string{url, property, name, content, string(time.Now().Unix())}
			writer.Write(information)
		})
	})
	writer.Flush()
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
