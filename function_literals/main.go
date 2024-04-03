package main

import (
	"fmt"
)

func main() {

	add := func(x int, y int) int {
		return x + y
	}

	multiply := func(x int, y int) int {
		return x * y
	}

	a, m := calculator(3, 5, add, multiply)

	fmt.Println(a, m)

}

func calculator(x int, y int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
