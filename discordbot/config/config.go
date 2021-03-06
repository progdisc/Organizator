package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// public variables
var (
	Token     string
	BotPrefix string

	// Private Variables
	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

// ReadConfig reads the configuration from a json file
func ReadConfig() error {
	fmt.Println("Reading from config file...")

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}
