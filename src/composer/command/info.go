package command

import (
	"composer/pckg"
)

type Info struct {
	Package string `long:"package" description:"Get info about this package" required:"true"`
}

func (self *Info) Execute(args []string) error {
	pckg := pckg.Pckg{self.Package, make([]*pckg.Version, 0)}
	pckg.Print()

	return nil
}
