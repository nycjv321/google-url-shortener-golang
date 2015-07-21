// google_url_shortener
package google_url_shortener

import (
	"errors"
	"fmt"
	"strings"
	"net/url"
	"net/http"
	"io"
	"encoding/json"
)

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

func (g * GoogleUrlShortener) ShortenedUrl(longUrl string) (string, error) {
	if strings.HasPrefix(longUrl, "https://goo.gl") && strings.HasPrefix(longUrl, "http://goo.gl")  {
		return "", errors.New("cannot shorten already shortened url")
	}
	
	client := &http.Client{}
	
	
	req, err := http.NewRequest("POST", g.postUrl(), strings.NewReader("{\"longUrl\": \"" + longUrl + "/\"}"))
	// ...
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	body_bytes:= resp.Body
	
	defer body_bytes.Close()
	
	var short_resp ShortenedResponse
	decoder := json.NewDecoder(body_bytes)
	err = decoder.Decode(&short_resp)
	if err != nil && err != io.EOF {
		return "", err
	}
	return short_resp.Id, err
}

func (g *GoogleUrlShortener) ExpandedUrl(id string) (string, error) {
	if !strings.HasPrefix(id, "https://goo.gl") && !strings.HasPrefix(id, "http://goo.gl")  {
		return "", errors.New("cannot expand non shortened url")
	}
	
	url, err := url.Parse(g.postUrl() + "&shortUrl=" + id)
	if err != nil {
		return "", err
	}
	
	client := &http.Client{}
	
	req, err := http.NewRequest("GET", url.String(), nil)
	// ...
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	body_bytes:= resp.Body
	
	defer body_bytes.Close()
	
	var exp_resp ExpandResponse
	decoder := json.NewDecoder(body_bytes)
	err = decoder.Decode(&exp_resp)
	if err != nil && err != io.EOF {
		return "", err
	}
	return exp_resp.LongUrl, err
}

type Response struct {
	Kind string `json:"kind"`
	Id string `json:"id"`
	LongUrl string `json:"longUrl"`
}

type ExpandResponse struct {
	Response
	Status string `json:"status"`
}

type ShortenedResponse struct {
	Response
}
