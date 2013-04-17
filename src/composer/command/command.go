package main

import (
	"composer/pckg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type composer struct {
	Packages map[string]map[string]*pckg.Pckg
}

func main() {
	src_json, e := ioutil.ReadFile("./Resources/packagist.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	//fmt.Printf("%s\n", string(src_json))

	u := composer{}
	err := json.Unmarshal(src_json, &u)
	if err != nil {
		panic(err)
	}

	for pckgName, pckg := range u.Packages {
		fmt.Printf("Project: %s\n", pckgName)
		for _, version := range pckg {
			version.Print()
			fmt.Printf("\n")
		}
	}

}
