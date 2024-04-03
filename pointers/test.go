package main

import (
	"fmt"
)

func Point() {

	var intVar int
	var pointerVar *int
	var pointerToPointerVar **int

	intVar = 100
	pointerVar = &intVar
	pointerToPointerVar = &pointerVar

	fmt.Println("intVar: \t\t\t", intVar)
	fmt.Println("pointerVar: \t\t", pointerVar)
	fmt.Println("pointerToPointerVar: \t", pointerToPointerVar)

	fmt.Println("")

	fmt.Println("&intVar: \t\t", &intVar)
	fmt.Println("&pointerVar: \t\t", &pointerVar)
	fmt.Println("&pointerToPointerVar: \t", &pointerToPointerVar)

	fmt.Println("")

	fmt.Println("*pointerVar: \t\t\t", *pointerVar)
	fmt.Println("*pointerToPointerVar: \t", *pointerToPointerVar)
	fmt.Println("**pointerToPointerVar: \t\t", **pointerToPointerVar)
}

func main() {
	Point()
}
