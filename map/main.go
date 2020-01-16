package main

import "fmt"

func main() {
	/*
		General syntax to define a Map:
		1. varname := map[<type-ofkey>] <type-of-value> { key1 : value1, ..., keyN : valueN }

		Different ways to use:
		1. var colors map[string]string OR colors := make{map[string]string}
			1.1. colors["green"] = "#008000"

	*/
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"blue":  "#0000FF",
	}
	fmt.Println(colors)
	delete(colors, "red")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k + " -> " + v)
	}
}
