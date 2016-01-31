package main
import (
	"os"
	"io/ioutil"
	"fmt"
	"bufio"
	"log"
)

func main() {
	//fileDemoReadWholeFile()
	//fileDemoReadBuffered()
	fileDemoReadBytes()
	// TODO: binary()
	// TODO: CSV file
}

func fileDemoReadWholeFile() {
	file, err := os.Create("test1.demo")

	if (err != nil) {
		panic(err)
	}
	file.WriteString("this is demo text")
	file.Close()

	read, _ := ioutil.ReadFile("test1.demo")
	fmt.Println(string(read))
}

func fileDemoReadBytes() {
	file, err := os.Open("bayareabikeshare/201508_trip_data-small.csv")

	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	file.Close()
}

func fileDemoReadBuffered() {
	file, err := os.Open("bayareabikeshare/201508_trip_data-small.csv")

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	//scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error while reading :", err)
	}

	fmt.Print(lines)
}
