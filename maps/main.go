package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
		"white": "#ffffff",
	}

	/*
		colors["white"] = "#ffffff"
		// map keys have to accessed with square brackets because map keys are typed
		fmt.Println(colors["white"])

		delete(colors, "white")
	*/
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is ", hex)
	}
}
