package main

import (
	"gopkg.in/yaml.v2"
	"github.com/hashicorp/consul/api"
	"io/ioutil"
)

type Configuration struct {
	Consul api.Config
	Port int
	Bind string
}

func LoadConfiguration(file string) (Configuration) {
	config := Configuration{}
	
	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		panic(err)
	}
	
	config.defaultValues()
	
	return config
}

func (c *Configuration) defaultValues() {
	if c.Bind == "" {
		c.Bind = "0.0.0.0"
	}
	
	if c.Port == 0 {
		c.Port = 8080
	}
}