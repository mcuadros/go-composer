package pckg

import (
	"fmt"
)

type Source struct {
	Type      string
	Url       string
	Reference string
}

func (s Source) Print() {
	fmt.Printf("\tType: %s, URL: %s\n", s.Type, s.Url)

}
