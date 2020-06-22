package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:dbname`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"database"`
}

// GetConfig load configuration from yaml
func GetConfig() Config {
	var configFile *os.File
	if _, err := os.Stat("configs/config.local.yaml"); os.IsNotExist(err) {
		configFile, _ = os.Open("configs/config.yaml")
	} else {
		configFile, err = os.Open("configs/config.local.yaml")
		if err != nil {
			panic(err)
		}
	}
	defer configFile.Close()

	var cnf Config
	decoder := yaml.NewDecoder(configFile)
	err := decoder.Decode(&cnf)
	if err != nil {
		panic(err)
	}

	log.Println("Config load successful")

	return cnf
}
