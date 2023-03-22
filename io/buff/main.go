package main

import (
	"bytes"
	"fmt"
	"strings"
)

//https://zenn.dev/hsaki/books/golang-io-package/viewer/bufio
func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())

	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}
