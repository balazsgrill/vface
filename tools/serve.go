package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fileServer)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
