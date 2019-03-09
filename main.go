package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("url")
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Error get" + url + "is fail")
	}

	fmt.Printf("[status] %d \n", resp.StatusCode)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[body] " + string(body))
}
