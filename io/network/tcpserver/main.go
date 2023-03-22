package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("can not Listen", err)
	}

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("can not accept", err)
	}

	responseData := []byte("hello net pkg!")
	_, err = conn.Write(responseData)
	if err != nil {
		fmt.Println("can not write", err)
	}
}
