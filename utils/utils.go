package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Config is a struct for backend config
type Config struct {
	Database struct {
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"password"`
		Port     string `json:"port"`
		Database string `json:"database"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

// LoadConfiguration is a load cfg from file
func LoadConfiguration() Config {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
	path = strings.Replace(path, "test", "", -1) + "/cfg.json"

	var config Config
	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

// AppConfig is a value loaded from config file
var AppConfig = LoadConfiguration()
