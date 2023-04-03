package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	ProxyOption string `json:"PROXY_OPTION"`
	APIKey      string `json:"APIKEY"`
	APIURL      string `json:"API_URL"`
	Model       string `json:"API_MODEL"`
	MaxTokens   int    `json:"API_MAX_TOKENS"`
	Temperature int    `json:"API_TEMPERATURE"`
	Stream      bool   `json:"API_STREAM"`
	APITimeout  int    `json:"API_TIMEOUT"`
}

func LoadConfig(filename string) (Config, error) {
	configFile, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()

	configData, err := ioutil.ReadAll(configFile)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
