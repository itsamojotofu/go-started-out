package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fileName := os.Args[1]
	fmt.Println("Trying to print out ", fileName)

	time.Sleep(500 * time.Millisecond)

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
