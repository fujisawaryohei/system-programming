package main

import (
	"io"
	"net"
	"os"
)

/// https://zenn.dev/hsaki/books/golang-io-package/viewer/netconn
func main() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")
	io.Copy(os.Stdout, conn)
}
