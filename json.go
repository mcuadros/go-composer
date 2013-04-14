package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
)

type composer struct {
    Age int
    Name string
    Homepage string
    Type string
    Require map[string]string
    Repositories []repositories
}

type repositories struct {
    Type string
    Url string
}

func main() {
    src_json, e := ioutil.ReadFile("./Resources/composer.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    fmt.Printf("%s\n", string(src_json))

    u := composer{}
    err := json.Unmarshal(src_json, &u)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("PHP:: %s\n", u.Require["php"])
    fmt.Printf("Project: %v\n", u.Name)
}