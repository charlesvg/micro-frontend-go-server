package internal

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"net/url"
	"os"
)

type AppConfig struct {
	Server struct {
		Port int `yaml:"port"`
		ContextPath string `yaml:"contextPath"`
	} `yaml:"server"`
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
}


type ProxyConfig struct {
	Server struct {
		Port int `yaml:"port"`
		DownStreamURL YAMLURL `yaml:"downStreamURL"`
	} `yaml:"server"`
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
}


func ReadAppConfig() AppConfig {
	f, err := os.Open(fmt.Sprintf("./configs/app.yml"))
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	var cfg = AppConfig{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Panicln(err)
	}
	return cfg
}

func ReadProxyConfig() ProxyConfig {
	f, err := os.Open(fmt.Sprintf("./configs/proxy.yml"))
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	var cfg = ProxyConfig{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Panicln(err)
	}
	return cfg
}

type YAMLURL struct {
	*url.URL
}

func (j *YAMLURL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}
	url, err := url.Parse(s)
	j.URL = url
	return err
}