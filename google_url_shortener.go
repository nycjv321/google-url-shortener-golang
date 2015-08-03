// google_url_shortener
package google_url_shortener

import (
	"errors"
	"fmt"
	"strings"
	"net/url"
	"io"
	"encoding/json"
	"http_utils"
)

// GoogleUrlShortener provides the ability to shorten links to goo.gl links. 
// It can also "unshorten" goo.gl links.
type GoogleUrlShortener struct {
	apiKey string
}

func (g *GoogleUrlShortener) postUrl() string {
	config, err := load_config()
	if err != nil {
		panic("Could not load config")
	}
	api_key := config.API_Key
	if api_key == "" {
		panic("Please define the api key used to authenticate requests")
	}

	return fmt.Sprintf("%v%v", "https://www.googleapis.com/urlshortener/v1/url?key=", api_key)
}

// Shorten a url using Google's goo.gl service
func (g * GoogleUrlShortener) ShortenedUrl(longUrl string) (string, error) {
	if strings.HasPrefix(longUrl, "https://goo.gl") && strings.HasPrefix(longUrl, "http://goo.gl")  {
		return "", errors.New("cannot shorten already shortened url")
	}
	
	body_reader := http_utils.Post(g.postUrl(), strings.NewReader("{\"longUrl\": \"" + longUrl + "/\"}"), "application/json")
	defer body_reader.Close()

	short_resp := deJsonifyShortenedResponse(body_reader)
	return short_resp.Id, nil
}

// Expand a goo.gl link
func (g *GoogleUrlShortener) ExpandedUrl(id string) (string, error) {
	if !strings.HasPrefix(id, "https://goo.gl") && !strings.HasPrefix(id, "http://goo.gl")  {
		return "", errors.New("cannot expand non shortened url")
	}
	
	url, err := url.Parse(g.postUrl() + "&shortUrl=" + id)
	if err != nil {
		return "", err
	}
	body_reader := http_utils.Get(url.String())
	defer body_reader.Close()

	exp_resp := deJsonifyExpendedResponse(body_reader)
	return exp_resp.LongUrl, nil
}

func deJsonifyExpendedResponse(reader io.Reader) (expand_response) {
	var r expand_response

	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&r)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return r
}

func deJsonifyShortenedResponse(reader io.Reader) (shortened_response) {
	var r shortened_response

	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&r)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return r
}

type response struct {
	Kind string `json:"kind"`
	Id string `json:"id"`
	LongUrl string `json:"longUrl"`
}

type expand_response struct {
	response
	Status string `json:"status"`

}

type shortened_response struct {
	response
}
