package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

//Config is a server's configuration
type Config struct {
	PORT       int    `json:"port"`
	StaticPath string `json:"server_static"`
}

func readConfig(file io.Reader) (Config, error) {
	conf := Config{}
	err := json.NewDecoder(file).Decode(&conf)
	return conf, err
}

func ReadConfigFile(filename string) (Config, error) {
	if !strings.HasSuffix(filename, ".json") {
		return Config{}, errors.New("JSON confix expected")
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		str := fmt.Sprintf("Config does not exist on (%s)", filename)
		return Config{}, errors.New(str)
	}
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}

	return readConfig(file)
}
