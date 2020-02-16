package main

import (
	"encoding/json"
	"os"
)

// Config ...
type Config struct {
	Token   string
	AdminID string
	Prefix  string
}

// Configure will try to read the config file
func Configure() Config {
	file, _ := os.Open("config.json")
	conf := Config{}
	json.NewDecoder(file).Decode(&conf)
	file.Close()
	return conf
}
