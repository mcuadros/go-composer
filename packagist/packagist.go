package packagist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gregjones/httpcache"
)

type packageResponse struct {
	Packages map[string]Packages `json:"packages"`
}

type Packagist struct {
	client  *http.Client
	current int
}

func NewPackagist() *Packagist {
	cache := httpcache.NewMemoryCache()
	transport := httpcache.NewTransport(cache)
	transport.Transport = &responseModifier{}

	return &Packagist{client: transport.Client()}
}

func (self *Packagist) GetPackages(name string) (Packages, error) {
	var response packageResponse

	body, err := self.doHTTPRequest(self.buildURL(name))
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &response)

	return response.Packages[name], nil
}

func (p *Packagist) buildURL(name string) string {
	return fmt.Sprintf("http://packagist.org/p/%s.json", name)
}

func (self *Packagist) doHTTPRequest(url string) ([]byte, error) {
	resp, err := self.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type responseModifier struct {
	Transport http.RoundTripper
}

func (t *responseModifier) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := t.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	resp, err := transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	resp.Header.Set("cache-control", "max-age=2592000")
	return resp, nil
}
