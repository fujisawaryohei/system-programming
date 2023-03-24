package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}

	// DOC: MultiWriter 可変長引数として複数のWriterを受け取り1つのWriterを返す
	writer := io.MultiWriter(file, os.Stdout)
	//　DOC: WriteString 対象のWriterに対して書き込みを行う
	io.WriteString(writer, "io.MultiWriter example\n")
}
