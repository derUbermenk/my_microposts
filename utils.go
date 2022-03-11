package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

// define a var config to be available
// 		for whole main package
var config Configuration

// init run before main package's main() function
func init() {
	// load config file
	loadConfig()
}

// parse config.json values
func loadConfig() {
	file, err := os.Open(`config.json`)

	// error should be nil to work
	// if it is not nil panic
	if err != nil {
		panic(err) // stop app
	}

	decoder := json.NewDecoder(file)

	// set initialize config
	config = Configuration{}

	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
}
