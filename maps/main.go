package main

import "fmt"

func main() {
	/*var names = []int{1, 2, 3}
	fmt.Print(names)
	*/

	// map declaration
	/*
		var names map[string]int
		names = make(map[string]int, 0)

		names["Muhammet"] = 1
		names["Can"] = 2
		names["Şimşek"] = 3

		fmt.Println(names) */

	names := map[string]int{
		"Muhammet": 1,
		"Can":      2,
		"Şimşek":   3,
	}
	fmt.Println(names)

	delete(names, "Can")

	fmt.Println(names)
}
