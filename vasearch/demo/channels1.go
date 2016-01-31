package main

import (
	"fmt"
	"time"
)

func get(cs <- chan string) string {
	var r string

	r = <-cs
	fmt.Println("received from channel: ", r)
	return r
}

func put(cs chan <- string) {
	val := "golang rules, java drools"
	cs <- val
}

func main() {
	c := make(chan string)
	go put(c)
	go get(c)

	time.Sleep(1 * 1e9)
	close(c)
}