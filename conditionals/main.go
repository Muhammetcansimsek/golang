package main

import "fmt"

func main() {
	var age = 18

	if age >= 18 {
		fmt.Print("Oy kullanabilirsiniz!")
	} else {
		fmt.Print("Oy kullanmak için 18 yaşının üzerinde olmalısınız!")
	}
}
