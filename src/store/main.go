package main

import (
	"log"
	"net/http"
	"os"
	"store/objects"
)


// starts with: LISTENER_ADDRESS=:12345 STORAGE_ROOT=./tmp go run main.go
func main() {
	http.HandleFunc("/objects/",objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTENER_ADDRESS"),nil))
}
