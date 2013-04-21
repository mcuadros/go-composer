package pckg

import (
	"fmt"
	"os"
	"strings"
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

func (self *Version) GetDependencies() []*Version {
	requires := make([]*Version, 0)
	//fmt.Printf("Looking for dependencies of %s\n", self.Name)

	for require, version := range self.Require {
		if require == "php" {
			continue
		}

		if strings.HasPrefix(require, "ext-") {
			continue
		}

		version, err := NewPckg(require).GetVersion(version)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}

		version.setRequiredBy(self)
		requires = append(requires, version.GetDependencies()...)
		requires = append(requires, version)
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
