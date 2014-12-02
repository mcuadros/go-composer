package pckg

import (
	"fmt"
)

type Source struct {
	Type      string
	Url       string
	Reference string
}

func (self *Source) Print() {
	fmt.Printf("\tType: %s, URL: %s\n", self.Type, self.Url)
}
