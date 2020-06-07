package internal

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
		ContextPath string `yaml:"contextPath"`
	} `yaml:"server"`
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
}

func ReadConfig() Config {
	const ConfigFileName = "./configs/application.yml"
	f, err := os.Open(ConfigFileName)
	if err != nil {
		log.Panicln("Unable to read config file", ConfigFileName, err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Panicln("Invalid yml config file", ConfigFileName, err)
	}
	return cfg
}
