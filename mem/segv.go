package main

import "fmt"

func main() {
	var p *int = nil

	fmt.Println("before injustice memory access")
	*p = 0
	fmt.Println("after injustice memory access")
}
