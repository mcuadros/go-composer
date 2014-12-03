package command

import (
	"github.com/mcuadros/go-composer/misc"
	"github.com/mcuadros/go-composer/packagist"
)

type Info struct {
	Verbosity []bool `short:"v" long:"verbose" description:"Show verbose debug information" `
	Package   string `short:"p" long:"package" description:"Get info about this package" required:"true"`
	Version   string `short:"s" long:"version" description:"Get info about this version" required:"true"`
}

func (self *Info) Execute(args []string) error {
	r := packagist.NewResolver()
	err := r.Resolve(self.Package, self.Version)
	if err != nil {
		panic(err)
	}

	for _, p := range r.Packages {
		misc.GetOutput().Info(p.String())
	}

	if len(self.Verbosity) == 0 {
		misc.GetOutput().SetVerbosity(misc.NOTICE)
	}

	return nil
}
