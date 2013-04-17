package pckg

import (
	"fmt"
)

type Pckg struct {
	Version  string
	Require  map[string]string
	Autoload *Autoload
	Source   *Source
}

type Autoload struct {
	PSR0     map[string]string
	ClassMap map[string]string
	Files    map[string]string
}

func (p Pckg) Print() {
	fmt.Printf("Version: %s, Require: %s\n", p.Version, p.Require)
	fmt.Printf("Source:\n")
	p.Source.Print()
}
