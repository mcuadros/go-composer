package main

import (
	"composer/pckg"
	"fmt"
)

func main() {
	pckg := pckg.Pckg{"yunait/mandango", make([]*pckg.Version, 0)}
	pckg.Print()

}
