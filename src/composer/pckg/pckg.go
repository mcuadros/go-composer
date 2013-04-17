package pckg

import (
	"composer/net"
	"encoding/json"
	"fmt"
)

type Pckg struct {
	Name     string
	Versions []*Version
}

func (self *Pckg) Print() {
	self.request()

	fmt.Printf("Package: %s\n", self.Name)
	for _, version := range self.Versions {
		version.Print()
	}
}

func (self *Pckg) request() bool {
	packagist := net.Packagist{}
	versions := packagist.GetRawVersion(self.Name)

	for number, raw := range versions {
		version := Version{}
		err := json.Unmarshal(raw, &version)
		if err != nil {
			return false
		}

		self.addVersion(&version)
		version.Number = number
	}

	return true
}

func (self *Pckg) addVersion(version *Version) {
	self.Versions = append(self.Versions, version)
}
