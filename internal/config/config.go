package config

import (
	"github.com/go-yaml/yaml"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	DBConfig struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSL      string `yaml:"sslmode"`
	} `yaml:"dbconfig"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("C:/Users/LEGION/Desktop/aitu/go/PostCreator/config.yml")
	if err != nil {
		log.Fatal("Error in reading file", err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return &config, nil
}
