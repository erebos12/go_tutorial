package main

import "fmt"

func main() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	val, ok := m["hello"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("No value found")
	}

	val, ok = m["none"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("No value found")
	}

}
