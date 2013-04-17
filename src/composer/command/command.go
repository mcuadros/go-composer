package main

import (
	"composer/net"
	"fmt"
	"os"
)

func main() {
	pckgName := os.Args[1]
	fmt.Printf("Project: %s\n", pckgName)

	packisg := net.Packagist{}
	pckg := packisg.GetPackage(pckgName)

	for _, version := range pckg {
		version.Print()
		fmt.Printf("\n")
	}
}
