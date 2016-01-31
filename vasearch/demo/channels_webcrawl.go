package main

import (
	"fmt"
	"io"
	"net/http"
	"io/ioutil"
	"time"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

var urls = []Lang{{Name:"golang", URL: "http://golang.org"}, {Name:"python", URL:"http://www.python.org"}}

func main() {
	//webcrawl()
	webcrawlWithTimeout()
}

func webcrawl() {
	start := time.Now()
	ch := make(chan string)
	n := len(urls)

	fmt.Println("num urls=", n)

	for _, url := range(urls) {
		fmt.Println(url)
		go count(url.Name, url.URL, ch)
	}

	for i := 0; i < n; i++ {
		fmt.Print(<-ch)
	}

	close(ch)

	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}

func count(name, url string, c chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		c <- fmt.Sprintf("%s %s %s", name, url, err)
	}

	n, _ := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	dt := time.Since(start).Seconds()

	c <- fmt.Sprintf("%s %d [%.2fs]\n", name, n, dt)
}

func webcrawlWithTimeout() {
	start := time.Now()
	ch := make(chan string)
	n := len(urls)

	fmt.Println("num urls=", n)

	for _, url := range (urls) {
		fmt.Println(url)
		go count(url.Name, url.URL, ch)
	}

	x := 0
	timeout := time.After(1 * time.Second)
	for {
		select {
		case result := <-ch:
			fmt.Print(result)

			x++
			if (x == n) {
				fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
				return
			}
		case <-time.After(10 * time.Millisecond):
			fmt.Printf(".")
		case <-timeout:
			fmt.Printf("timed out")
			return
		}
	}

	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}



