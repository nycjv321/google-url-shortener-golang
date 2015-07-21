package google_url_shortener

import (
	"testing"
)

func TestLoad(t *testing.T) {
	g, error := load_config()
	if error != nil {
		panic("Could not load config. config is incorrect: " + g.API_Key)
	}
}
