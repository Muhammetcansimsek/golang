package main

import "fmt"

//A slice have been defined
var numbers = []int{10, 20, 2, 3, 5}

//Function sumSlice() has been created in order to sum slice named 'numbers'
func sumSlice(numbers []int) int {
	sum := 0

	//for-each has been created by leaving the first argument blank
	for _, value := range numbers {
		sum += value
	}
	return sum
}

//Function add(), takes two argument named x and y. Then sums these argument and then returns.
func add(x int32, y int32) (string, int32) {
	return "Sum of The Numbers: ", (x + y)
}

func main() {
	fmt.Print(add(10, 20))
}
