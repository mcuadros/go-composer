package pckg

import (
	"fmt"
)

type Version struct {
	Number   string
	Name     string
	Require  map[string]string
	Autoload *Autoload
	Source   *Source
}

type Autoload struct {
	PSR0     map[string]string
	ClassMap map[string]string
	Files    map[string]string
}

func (self *Version) Print() {
	fmt.Printf("Name: %s, Number: %s\n", self.Name, self.Number)
	fmt.Printf("Require: %s\n", self.Require)
	fmt.Printf("Source:\n")
	self.Source.Print()

	fmt.Printf("\n")
}
