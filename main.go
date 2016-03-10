package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
)

var config Configuration

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("fatal error: %s", r)
		}
	}()
	
	config = LoadConfiguration()
	config.Check()

	http.HandleFunc("/", forwardHandler)
	
	config.Print()
	http.ListenAndServe(config.Bind, nil)
}

func (c Configuration) Print() {
	log.Printf("http: http://%s\n", config.Bind)
	log.Printf("Consul:\n")
	log.Printf("  Address: %s\n", c.Consul.Address)
	log.Printf("  Datacenter: %s\n", c.Consul.Datacenter)
}

func (c Configuration) Check() {
	message := ""
	
	if c.Bind == "" {
		message += "missing bind address\n"
	}
	
	if c.Consul.Address == "" {
		message += "consul server address missing\n"
	}
	
	if message != "" {
		fmt.Printf("error:\n%s", message)
		os.Exit(1)
	}
}