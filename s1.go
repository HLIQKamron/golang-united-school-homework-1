package io_reader

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func ReaderSplit(strReader *strings.Reader, n int) []string {
	buffer := make([]byte, 2)
	str := ""
	for {
		_, err := strReader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		if string(buffer) != "\n" && string(buffer) != "\t" {
			str += string(buffer)
		}
	}
	str = strings.Join(strings.Split(strings.Join(strings.Split(str, "\n"), ""), "\t"), " ")
	str2 := ""
	arr := []string{}
	for i := 0; i < len(str); i++ {
		if string(str[i]) != " " && string(str[i]) != "\t" {
			str2 += string(str[i])
		}
		if len(str2) == n || i == len(str)-1 {
			arr = append(arr, str2)
			str2 = ""
		}

	}
	return arr
}
func SeekTillHalfOfString(strReader *strings.Reader, n int) []string {
	buf := make([]byte, strReader.Len())
	nn, err := strReader.Read(buf)
	if err == io.EOF {
		log.Println("we reached EOF")
	} else if err != nil {
		log.Fatal(err)
	}

	return []string{string(buf[nn/2:])}
}
