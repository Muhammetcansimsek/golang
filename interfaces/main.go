package main

/*
*Aslında buradaki mantık şu ki, örneğin benim bir mağazam var ve birden fazla kategoride ürün satışı yapıyorum ve
* bu ürünlerden karışık bir şekilde ve istediği kadarını sepete ekleyen kullanıcının sepete gittiğinde total tutarı görmesi için
* aslında her kategori için ayrı bir total tutar hesaplayan bir fonksiyon yazmak yerine bunları interface ile soyutlayarak ve bu
* interface'i implement eden structlar ile tek bir calculateTotalShippingCost() isimli fonksiyonla tüm kategorilerde toplam sepet tutarını
* hesaplayan bir fonksiyon/metot tanımlamak dinamiklik ve kolaylık sağlayacağı için burada yapılan budur aslında tamamen 'soyutlama',
* 'arayüz tanımlama' ve 'polymorphism' kullanılarak yapılan bir işlem
 */

//Interface has been declarated
type IShippable interface {
	ShippingCost() int
}

type Flower struct {
	desi int
}

func (flower *Flower) ShippingCost() int {
	return 5 + flower.desi*3
}

type Book struct {
	desi int
}

//Implemented, interface's function shippinCoast()
func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}

type Electronic struct {
	desi int
}

//Implemented, interface's function shippinCoast()
func (electronics *Electronic) ShippingCost() int {
	return 5 + electronics.desi*4
}

func calculateTotalShippingCost(products []IShippable) int {
	total := 0
	for _, product := range products {
		total += product.ShippingCost()
	}
	return total
}

/*
func calculateTotalShippingCostofElectronics(electronics []Electronic) {
	total := 0
	for _, electronic := range electronics {
		total += electronic.ShippingCoast()
	}
}

func calculateTotalShippingCostofBooks(books []Book) {
	total := 0
	for _, book := range books {
		total += book.ShippingCoast()
	}
}
*/

/*
func main() {
	var products []IShippable = []IShippable{
		&Book{desi: 10},
		&Electronic{desi: 8},
		&Flower{desi: 12},
	}
	fmt.Println(calculateTotalShippingCost(products))

	/*
		book1 := &Book{desi: 10}
		book2 := &Book{desi: 20}

		fmt.Println(book1.ShippingCoast())
		fmt.Println(book2.ShippingCoast())

}
*/
