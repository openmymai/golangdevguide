package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	// colors := make(map[int]string)
	// colors[10] = "#ffffff"

	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}