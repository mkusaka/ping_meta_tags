package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func getUrls() []string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	urlstrings := os.Getenv("url")
	urls := strings.Split(urlstrings, ",")

	return urls
}

func scrape(urls []string) {

	file, err := os.Create("tmp/result.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	for idx, url := range urls {
		doc, err := goquery.NewDocument(url)

		if err != nil {
			log.Fatal(err)
		}

		headers := []string{"url", "property", "name", "content", "timestamp"}
		writer.Write(headers)
		doc.Find("head").Each(func(i int, s *goquery.Selection) {
			s.Find("meta").Each(func(j int, m *goquery.Selection) {
				property, _ := m.Attr("property")
				name, _ := m.Attr("name")
				content, _ := m.Attr("content")
				information := []string{urls[idx], property, name, content, string(time.Now().Unix())}
				writer.Write(information)
			})
		})
	}
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
	urls := getUrls()
	scrape(urls)
	fmt.Println("finish scrape" + urls)
}
