package main

import "fmt"

type ten []int

func main() {
	t := ten{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range t {
		if i%2 != 0 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}
	}
}
