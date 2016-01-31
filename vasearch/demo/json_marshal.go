package main

import (
	"encoding/json"
	"fmt"
	"encoding/xml"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func main() {
	lang := Lang{"Go", 2009, "http://golang.org"}
	fmt.Printf("%v %v+ %#v\n", lang)

	data, err := json.MarshalIndent(lang, "", " ")

	if (err != nil) {
		panic(err)
	}
	fmt.Println(string(data))

	data, _ = xml.MarshalIndent(lang, "", " ")

	fmt.Println(string(data))
}