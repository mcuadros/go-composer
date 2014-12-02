package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcuadros/go-composer/misc"
)

type Packagist struct {
	Name     string
	Raw      []byte
	Response struct {
		Packages map[string]map[string]json.RawMessage
	}
}

func (self *Packagist) request() bool {
	base := "http://packagist.org/p/%s.json"
	url := fmt.Sprintf(base, self.Name)

	res, err := http.Get(url)
	if err != nil {
		return false
	}

	raw, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return false
	}

	self.Raw = raw
	misc.GetOutput().Debug("Url: %s (%d bytes)", url, len(raw))

	return true
}

func (self *Packagist) unmarshal() bool {
	err := json.Unmarshal(self.Raw, &self.Response)
	if err != nil {
		return false
	}

	return true
}

func (self *Packagist) GetRawVersion(pckgName string) map[string]json.RawMessage {
	self.Name = pckgName

	self.request()
	self.unmarshal()

	return self.Response.Packages[self.Name]
}
