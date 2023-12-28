package main

import "fmt"

func main()  {

	// var colors map[string] string

	// colors := make(map[int]string)

	colors := map[string] string {
		"red": "#ff0000",
		"green": "#00FF7F",
		"white": "#fff",
	}

	printMap(colors)
}

func printMap(c map[string]string){
	for color, hex := range c {
		fmt.Println("Hexx code for", color, "is", hex)
	}
}