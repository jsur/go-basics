package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)

	/*
		colors = map[string]string{
			"red":   "#ff0000",
			"green": "#00FF00",
		}
	*/

	colors["white"] = "#ffffff"

	// map keys have to accessed with square brackets because map keys are typed
	fmt.Println(colors["white"])

	delete(colors, "white")

	fmt.Println(colors)
}
