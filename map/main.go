package main

import "fmt"

func main() {
	//var colors map[string]string

	colors := make(map[string]string)

	//colors := map[string]string{
	//	"red":   "#ff0000",
	//	"green": "#dd2402",
	//	"blue":  "#ld1122",
	//}

	colors["red"] = "#ff0000"
	colors["green"] = "#dd2402"

	delete(colors, "green")

	fmt.Println(colors)
}
