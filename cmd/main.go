package main

import (
	io_reader "golang-united-school-homework-1"
	"strings"
)

func main() {
	someString := "HELLOWORLD"
	myReader := strings.NewReader(someString)
	io_reader.SeekTillHalfOfString(myReader, 5)
}
