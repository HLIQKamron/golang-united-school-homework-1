package io_reader

import (
	"fmt"
	"io"
	"strings"
)

func ReaderSplit(strReader *strings.Reader, n int) []string {
	buffer := make([]byte, n)
	str := ""
	arr := []string{}
	for {
		_, err := strReader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		str += string(buffer)
		arr = append(arr, str)
		str = ""
	}
	return arr
}
func SeekTillHalfOfString(strReader *strings.Reader, n int) []string {
	buf := make([]byte, strReader.Len())
	nn, _ := strReader.Read(buf)

	return []string{string(buf[nn/2:])}
}
