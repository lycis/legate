package main

import (
	"log"
	"net/http"
	"fmt"
)

var config Configuration

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("fatal error: %s", r)
		}
	}()

	config = LoadConfiguration("legate.yml")

	http.HandleFunc("/", forwardHandler)
	
	config.Print()
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}

func (c Configuration) Print() {
	log.Printf("Listening: http://%s:%d\n", config.Bind, config.Port)
	log.Printf("Consul:\n")
	log.Printf("  Address: %s\n", c.Consul.Address)
	log.Printf("  Datacenter: %s\n", c.Consul.Datacenter)
}