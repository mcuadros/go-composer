package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "log"
)

type composer struct {
    Age int
    Name string
    Homepage string
    Type string
    Packages map[string]test
    Repositories []test
}

type test struct {
    name string
    Url string
}

func main() {
    res, err := http.Get("http://packagist.org/p/yunait/mandango.json");
    if err != nil {
        log.Fatal(err)
    }

    src_json, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s\n", string(src_json))


    u := composer{}
    error := json.Unmarshal(src_json, &u)
    if error != nil {
        panic(error)
    }
    
    fmt.Printf("PHP:: %s\n", u.Packages["yunait/mandango"])
    fmt.Printf("Project: %v\n", u.Name)
}