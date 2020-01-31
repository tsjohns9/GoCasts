package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// colors := make(map[string]string)

	colors["blue"] = "blue"

	delete(colors, "white")

	primtMap(colors)
}

func primtMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
