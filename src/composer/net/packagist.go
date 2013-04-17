package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	res, err := http.Get(fmt.Sprintf(base, self.Name))
	if err != nil {
		return false
	}

	raw, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return false
	}

	self.Raw = raw

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
