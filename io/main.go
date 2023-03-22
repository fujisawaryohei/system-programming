package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
