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

func touchFile(filename string) {
	if fileOrDirExistence(filename) {
		return
	}

	file, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func makeTmpDir() error {
	if fileOrDirExistence("tmp") {
		return nil
	}

	return os.Mkdir("tmp", 0777)
}

func fileOrDirExistence(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func main() {
	err := makeTmpDir()

	if err != nil {
		log.Fatal("can't make tmp under current directory.")
	}

	touchFile(".env")

	urls := getUrls()

	if urls[0] == "" {
		log.Fatal("url must be at .env file (or enviroment variable). like `url=url1,url2` format.")
		return
	}

	scrape(urls)
	fmt.Println("finish scrape" + strings.Join(urls, ","))
}
