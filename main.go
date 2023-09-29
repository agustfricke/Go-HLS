package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type VideoServe struct {
}

func (h VideoServe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func main() {
	http.Handle("/yoyo", VideoServe{})
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Printf("Starting server on %v\n", 8080)
	log.Printf("Serving %s on HTTP port: %v\n", ".", 8080)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 8080), nil))
}
