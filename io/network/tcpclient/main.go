package main

import (
	"io"
	"net"
	"os"
)

/// https://zenn.dev/hsaki/books/golang-io-package/viewer/netconn
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: localhost\r\n\r\n")
	io.Copy(os.Stdout, conn)
}
