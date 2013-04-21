package command

import (
	"composer/misc"
	"composer/pckg"
)

type Info struct {
	Verbosity []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Package   string `short:"p" long:"package" description:"Get info about this package" required:"true"`
	Version   string `short:"s" long:"version" description:"Get info about this version"`
}

func (self *Info) Execute(args []string) error {
	pckg := pckg.NewPckg(self.Package)

	if len(self.Verbosity) == 0 {
		misc.GetOutput().SetVerbosity(misc.NOTICE)
	}

	if len(self.Version) != 0 {
		version, err := pckg.GetVersion(self.Version)
		if err != nil {
			return err
		}

		version.Print()
		return nil
	}

	pckg.Print()
	return nil

}
