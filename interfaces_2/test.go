package main

import "fmt"

type ISiparis interface {
	singleCost() int
}

type cicek struct {
	demet int
}

func (cicek *cicek) singleCost() int {
	return cicek.demet * 30
}

type cikolata struct {
	paket int
}

func (cikolata *cikolata) singleCost() int {
	return cikolata.paket * 50
}

func calculatetotalCost(products []ISiparis) float32 {
	total := 0
	for _, iter := range products {
		total += iter.singleCost()
	}
	return float32(total)
}

func main() {
	var products []ISiparis = []ISiparis{
		&cicek{demet: 15},
		&cikolata{paket: 5},
	}
	fmt.Println(calculatetotalCost(products))

}
