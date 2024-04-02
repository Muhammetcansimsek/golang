package main

import "fmt"

func main() {
	// We can create customers as many as we want that defined struct
	/*var customer1 = Customer{id: 1, name: "Muhammet Can Şimşek", age: 24,
	address: Address{city: "1111.sk no:1111, 1111Apart, Isparta,Türkiye", district: "Merkez"}}*/
	var customer2 = Customer{id: 2, name: "Mustafa Murat Coşkun", age: 30,
		address: Address{city: "1111.sk no:1111, 1111Apart, Ankara,Türkiye", district: "Kızılay"}}

	//fmt.Println(customer1)
	//fmt.Println(customer2)

	customer2.ToString()
	customer2.changeName("Mustafa Murat")
	customer2.ToString()
}

// struct has been created
type Customer struct {
	id      int32
	name    string
	age     uint16
	address Address
}

type Address struct {
	city     string
	district string
}

/*
func changeName(customer *Customer) {
	customer.name = "Murat Coşkun"
}
*/

func (customer Customer) ToString() {
	fmt.Printf("Name: %s, Age: %d\n", customer.name, customer.age)
}

func (customer *Customer) changeName(newName string) {
	customer.name = newName
}
