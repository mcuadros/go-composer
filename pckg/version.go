package pckg

import (
	"fmt"
	"strings"

	"github.com/mcuadros/go-composer/misc"
)

type Version struct {
	Name       string
	Version    string
	Require    map[string]string
	Autoload   *Autoload
	Source     *Source
	RequiredBy *Version
}

type Autoload struct {
	PSR0     map[string]string `json:"psr-0"`
	ClassMap []string
	Files    []string
}

func (self *Version) GoGetDependencies(c chan []*Version) {
	misc.GetOutput().Error("New GoRutine")

	c <- self.GetDependencies()
}

func (self *Version) GetDependencies() []*Version {
	requires := make([]*Version, 0)
	channels := make([]chan []*Version, 0)

	for require, version := range self.Require {
		if require == "php" {
			continue
		}

		if strings.HasPrefix(require, "ext-") {
			continue
		}

		if version == "self.version" {
			//misc.GetOutput().Error("Self.version: %s", self.Version)
			version = self.Version
		}

		misc.GetOutput().Info("%s\n", version)

		version, err := NewPckg(require).GetVersion(version)
		if err != nil {
			misc.GetOutput().Error(err.Error())
			return nil
		}

		version.setRequiredBy(self)
		channel := make(chan []*Version)
		channels = append(channels, channel)
		self.GoGetDependencies(channel)

		requires = append(requires, version)
	}

	for _, channel := range channels {
		misc.GetOutput().Error("Reading GoRutine")
		requires = append(requires, <-channel...)
	}

	return requires
}

func (self *Version) setRequiredBy(version *Version) {
	self.RequiredBy = version
}

func (self *Version) Print() {
	fmt.Printf("Name: %s, Number: %s\n", self.Name, self.Version)

	fmt.Printf("Require:\n")
	for index, version := range self.GetDependencies() {
		fmt.Printf("\t[%d] %-20s (%-20s)\t%s\n", index, version.Name, version.Version, version.RequiredBy.Name)
	}

	fmt.Printf("Source:\n")
	self.Source.Print()

	fmt.Printf("\n")
}
