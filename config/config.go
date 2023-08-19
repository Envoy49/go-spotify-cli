package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ServerUrl       string
	Port            string
	RequestedScopes string
	ClientId        string
	ClientSecret    string
}

var GlobalConfig Config

func LoadConfiguration() {

	file, err := os.Open("./config.json")
	if err != nil {
		fmt.Println(&err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
		return
	}
	GlobalConfig = Config{
		ServerUrl:       config.ServerUrl,
		Port:            config.Port,
		RequestedScopes: config.RequestedScopes,
		ClientId:        config.ClientId,
		ClientSecret:    config.ClientSecret,
	}
}
