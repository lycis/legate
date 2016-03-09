package main

import (
	"log"
	"net/http"
	"fmt"
)

var config Configuration

func main() {
	log.Println("Starting legate...")
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("fatal error: %s", r)
		}
	}()

	config = LoadConfiguration("legate.yml")

	http.HandleFunc("/", forwardHandler)
	
	log.Printf("Listening: http://%s:%d\n", config.Bind, config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
