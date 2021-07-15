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

	fmt.Println(colors)
}
