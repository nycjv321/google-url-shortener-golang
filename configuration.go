package google_url_shortener

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// configuration provides the data used by GoogleUrlShortener. 
type configuration struct {
	API_Key string
}

func open_error() error {
	return errors.New("Could not open config.json")
}

// Read the config via JSON. It looks for a config.json in the current directory.
func load_config() (*configuration, error) {
	config_file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, open_error()
	}
	c := &configuration{}
	json.Unmarshal(config_file, c)
	return c, nil
}
