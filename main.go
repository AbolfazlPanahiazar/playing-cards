package main

import "fmt"

func main() {
	// var colors map[string]string
	// or
	// colors := make(map[string]string)
	// or
	colors := map[string]string{
		"white": "#FFFFFF",
		"red":   "#FF0000",
	}

	colors["blue"] = "#0000FF"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
