package pckg

import (
	"composer/misc"
	"composer/net"
	"encoding/json"
	"fmt"
)

type Pckg struct {
	Name      string
	Versions  map[string]*Version
	requested bool
}

var cache = make(map[string]*Pckg, 0)

func NewPckg(name string) *Pckg {
	if pckg, ok := cache[name]; ok {
		return pckg
	}

	cache[name] = new(Pckg)
	cache[name].SetName(name)

	return cache[name]
}

func (self *Pckg) GetVersion(name string) (*Version, error) {
	self.request()

	if version, ok := self.Versions[name]; ok {
		return version, nil
	}

	misc.GetOutput().Warning("cannot find given version %s at %s", name, self.Name)

	for _, version := range self.Versions {
		return version, nil
	}

	return nil, fmt.Errorf("cannot find given version %s at %s", name, self.Name)
}

func (self *Pckg) GetName() string {
	return self.Name
}

func (self *Pckg) SetName(name string) {
	self.Name = name
}

func (self *Pckg) Print() {
	self.request()

	fmt.Printf("Package: %s\n", self.Name)
	for _, version := range self.Versions {
		fmt.Printf("\t%-20s (%-20s)\n", version.Name, version.Version)
	}
}

func (self *Pckg) request() bool {
	if self.requested {
		return true
	}

	packagist := net.Packagist{}
	versions := packagist.GetRawVersion(self.Name)

	for number, raw := range versions {
		version := Version{}
		err := json.Unmarshal(raw, &version)
		if err != nil {
			fmt.Printf("Version: %s Error:%s Data:%s\n", number, err)
		}

		self.addVersion(number, &version)

	}

	self.requested = true
	return true
}

func (self *Pckg) addVersion(number string, version *Version) {
	if self.Versions == nil {
		self.Versions = make(map[string]*Version)
	}

	self.Versions[number] = version
}
