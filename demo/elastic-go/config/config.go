package config

import (
	"io/ioutil"
	"log"
	
	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Elastic Elastic `json:"elastic" yaml:"elastic"`
}

type Elastic struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func LoadAndInit(conf string) *ServerConfig {
	file, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal("read yaml file failed")
	}
	
	sconf := &ServerConfig{}
	err = yaml.UnmarshalStrict(file, &sconf)
	if err != nil {
		log.Fatal("yaml unmarshal failed")
	}
	
	return sconf
}
