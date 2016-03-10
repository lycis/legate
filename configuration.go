package main

import (
	"gopkg.in/yaml.v2"
	"github.com/hashicorp/consul/api"
	"io/ioutil"
	"flag"
)

type Configuration struct {
	Consul api.Config
	Bind string
}

func LoadConfiguration() (Configuration) {
	config := Configuration{}
	
	// parse commandline first
	file := flag.String("config", "", "configuration file")
	datacenter := flag.String("dc", "dc01", "consul datacenter")
	bind := flag.String("bind", ":8080", "address and port to bind to (e.g. 127.0.0.1:80) ")
	
	flag.Parse()
	
	consulAddress := flag.Arg(0)
	
	if *file != "" {
		loadConfigFromFile(*file)
	} else {
		config.Bind = *bind
		config.Consul.Address = consulAddress
		config.Consul.Datacenter = *datacenter
	}
	
	return config
}

func loadConfigFromFile(file string){
	// read file
	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	
	err = yaml.Unmarshal(fileContent, &config)
		if err != nil {
			panic(err)
		}
		config.defaultValues()
		return
}

func (c *Configuration) defaultValues() {
	if c.Bind == "" {
		c.Bind = "0.0.0.0:8080"
	}
}