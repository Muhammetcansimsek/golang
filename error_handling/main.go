package main

import (
	"errors"
	"fmt"
)

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct {
}

func (productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return errors.New("Ürün ismi boş bırakılamaz!")
	}
	if product.price <= 0 {
		return errors.New("Ürün fiyatı sıfırdan büyük olmalı!")
	} else {
		fmt.Println("Ürün başarıyla eklendi.")
		return nil
	}
}

func main() {
	productService := ProductService{}

	err := productService.Add(Product{id: 1, name: "Çamaşır makinesi", price: -10000})

	if err != nil {
		fmt.Println(err)
	}
}

/*
func divide(pay int, payda int) (float32, error) {
	if payda == 0 {
		return 0, errors.New("Payda 0 olamaz..!")
	} else {
		return (float32(pay / payda)), nil
	}
}
*/

/*
	result, err := divide(10, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
*/
/*
	fileData, err := os.ReadFile("sample.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileData)
	}
*/
