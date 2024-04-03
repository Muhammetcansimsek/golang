package main

import (
	"fmt"
)

func main() {
	/*fmt.Println("Hello")
	fmt.Println("World")

	defer fmt.Println("Hello") // wait
	fmt.Println("Golang!")
	*/

	// defer anahtar kelimesi queue yapıdadır ve LIFO mantığıyla çalışır
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("Defer anahtar kelimesinin kullanımı")
}
