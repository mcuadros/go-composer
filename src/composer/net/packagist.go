package net

import (
	"composer/pckg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Packagist struct {
	Name     string
	Raw      []byte
	Response *response
}

type response struct {
	Packages map[string]map[string]*pckg.Pckg
}

func (p *Packagist) request() bool {
	base := "http://packagist.org/p/%s.json"
	fmt.Println(fmt.Sprintf(base, p.Name))

	res, err := http.Get(fmt.Sprintf(base, p.Name))
	if err != nil {
		return false
	}

	raw, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return false
	}

	p.Raw = raw

	return true
}

func (p *Packagist) unmarshal() bool {
	err := json.Unmarshal(p.Raw, &p.Response)
	if err != nil {
		return false
	}

	fmt.Println(p.Name)

	return true
}

func (p *Packagist) GetPackage(pckgName string) map[string]*pckg.Pckg {
	p.Name = pckgName

	p.request()
	p.unmarshal()

	return p.Response.Packages[p.Name]
}
