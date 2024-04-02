package main

import "fmt"

func main() {
	var a int  //Normal variable
	var p *int //Pointer has been declarated

	a = 10
	p = &a

	fmt.Println("Pointer's pointed address:", p)              //Pointer's pointed address
	fmt.Println("The value of pointer's pointed address", *p) //The value of pointer's pointed address

	*p = 20 //The value of pointer's pointed address has been changed

	fmt.Println(a)  //Relatively, the value of a variable has been changed
	fmt.Println(*p) //It has been mentioned previously

	var b int = 10
	add12(b)
	fmt.Println("The value of b before:", b)

	add12pointer(&b)
	fmt.Println("The value of b after:", b)

	var numbers = []int{1, 2, 3}
	fmt.Println(numbers)
	changeValue(numbers)
	fmt.Println(numbers)
}

func changeValue(numbers []int) {
	numbers[0] = 1000
}

func add12(x int) { // x == b
	x += 12 // x = x + 12, x => 22, b == 10
}

func add12pointer(x *int) {
	*x += 12 //The x points address of b variable, *x means value of the pointed address. *x == 22 and b == 22
}
