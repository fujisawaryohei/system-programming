package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func main() {
	pid := os.Getegid()
	fmt.Println("virtual address spaces before memory map")
	command := exec.Command("cat", "/proc/"+strconv.Itoa(pid)+"/maps")
	command.Stdout = os.Stdout
	err := command.Run()
	if err != nil {
		log.Fatal("failed to cat")
	}

	file, err := os.OpenFile("testfile", os.O_RDWR, 0)
	if err != nil {
		log.Fatal("can not open testfile")
	}
	defer file.Close()

	data, err := syscall.Mmap(int(file.Fd()), 0, 5, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatal("failed to mmap")
	}

	fmt.Println("")
	fmt.Printf("testfile memory map to: %p\n", &data[0])
	fmt.Println("")

	fmt.Println("virtual address spaces after memory map")
	command = exec.Command("cat", "/proc/"+strconv.Itoa(pid)+"/maps")
	command.Stdout = os.Stdout
	err = command.Run()
	if err != nil {
		log.Fatal("failed to cat")
	}

	replaceBytes := []byte("Hello")
	for i, _ := range data {
		data[i] = replaceBytes[i]
	}
}
