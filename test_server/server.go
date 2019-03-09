package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func handle(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	fmt.Println("[method] " + method)
	for k, v := range req.Header {
		fmt.Print("[header]" + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	if method == "GET" {
		req.ParseForm()
		params := ""
		for k, v := range req.Form {
			fmt.Print("[param] " + k)
			params += "[param] " + k
			fmt.Println(": " + strings.Join(v, ", "))
			params += ": " + strings.Join(v, ", ") + ","
		}
		t := template.Must(template.ParseFiles("test_server/template000.html.tpl"))
		if err := t.ExecuteTemplate(w, "template000.html.tpl", params); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":9999", nil)
}
