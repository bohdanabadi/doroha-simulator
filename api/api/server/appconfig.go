package server

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	ServerDev struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"serverDev"`
	ServerProd struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"serverProd"`
}

func LoadConfig() Config {
	var config Config
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("apperror: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("apperror: %v", err)
	}

	return config
}
