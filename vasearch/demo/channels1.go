package main

import (
	"fmt"
	"time"
	"runtime"
)

// START OMIT
func get(cs <- chan string)  {
	r := <-cs
	fmt.Println("received from channel: ", r)
}

func put(cs chan <- string) {
	val := "golang rules, java drools"
	cs <- val
}

func main() {
	fmt.Println("#cpus= ",  runtime.NumCPU())
	c := make(chan string)
	go put(c)
	go get(c)

	time.Sleep(1 * time.Second)
	close(c)
}
//END OMIT