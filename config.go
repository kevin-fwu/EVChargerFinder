package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Nrel struct {
		File  string
		Token string
	}
	Server struct {
		Address string
		Ssl     struct {
			Key  string
			Cert string
		}
	}
}

func LoadConf(filename string) (*Config, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
