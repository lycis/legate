package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"net/http"
	"strings"
)

var consulConfig *api.Config

// Forward Handler
// forwards the caller to the according service retrieved from consul
func forwardHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	firstSlashAt := strings.Index(path, "/")
	
	var serviceName string
	var servicePath string
	if firstSlashAt < 0 {
		serviceName = path
		servicePath = ""		
	} else {
		serviceName = path[0:firstSlashAt]
		servicePath = path[firstSlashAt:]
	}
	
	if len(serviceName) < 1 {
		http.NotFound(w, r)
		return
	}

	log.Printf("redirect request for service '%s', resource '%s'\n", serviceName, servicePath)

	defer func() {
		if r := recover(); r != nil {
			log.Fatal("recovered from error (%s). sending internal error\n", r)
			http.Error(w, r.(string), http.StatusInternalServerError)
			panic(r)
		}
	}()

	client, err := api.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}
	log.Printf("connection to consul established\n")

	catalog := client.Catalog()

	serviceList, _, err := catalog.Service(serviceName, "", nil)
	if err != nil {
		panic(err)
	}

	if len(serviceList) < 1 {
		log.Printf("service '%s' not found\n", serviceName)
		http.NotFound(w, r)
		return
	} else {
		service := serviceList[0]
		log.Printf("service %s (%s) identified with address '%s:%d'\n", service.ServiceName, service.ServiceID, service.ServiceAddress, service.ServicePort)
		serviceUrl := fmt.Sprintf("http://%s:%d%s", service.ServiceAddress, service.ServicePort, servicePath)
		log.Printf("fowarding to: %s\n", serviceUrl)
		http.Redirect(w, r, serviceUrl, http.StatusMovedPermanently)
	}

}

func loadConfiguration() {
	cconf := api.Config{Address: "vvieyggprod01:8500", Datacenter: "yg01"}

	consulConfig = &cconf
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	loadConfiguration()

	http.HandleFunc("/", forwardHandler)
	http.ListenAndServe(":8080", nil)
}
