package main

import "fmt"

func main() {
	/*
		// defined a slice which is integer
		var numbers = []int{1, 2, 3, 4}

		// for-each loop
		for _, value := range numbers {
			fmt.Println(value)
		}
	*/

	names := map[string]int{
		"muhammet": 17,
		"can":      30,
		"şimşek":   89,
	}

	for key, value := range names {
		fmt.Printf("Anahtar: %s , Değer: %d\n", key, value)
	}
}
