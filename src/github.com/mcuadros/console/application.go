package console

import (
	"flag"
	"github.com/mcuadros/console/output"
)

type Application struct {
	test string
}

func (self *Application) Run(output output.Output) {
	flag.StringVar(&self.test, "gopher_type", "default", "usage")
	output.Emergency(self.test)

	flag.Parse()
}
