package google_url_shortener

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Configuration struct {
	API_Key string
}

func openError() error {
	return errors.New("Could not open config.json")
}

func load_config() (*Configuration, error) {
	config_file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, openError()
	}
	c := &Configuration{}
	json.Unmarshal(config_file, c)
	return c, nil
}
