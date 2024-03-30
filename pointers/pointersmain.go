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

}
