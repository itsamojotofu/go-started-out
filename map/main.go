package main

import "fmt"

func main() {
	//var colors map[string]string
	//colors := make(map[string]string)
	colors := map[string]string{
		"blue":  "#ld1122",
		"white": "#ffffff",
	}

	colors["red"] = "#ff0000"
	colors["green"] = "#dd2402"

	delete(colors, "green")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
