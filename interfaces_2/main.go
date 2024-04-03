package main

import "fmt"

type IEncoder interface {
	Encode(value string)
	Decode(value string)
}

type yEncoder struct {
}

func (yencoder *yEncoder) Encode(value string) {
	fmt.Println("yEncoder ile kodlandı.")
}
func (yencoder *yEncoder) Decode(value string) {
	fmt.Println("yEncoder ile kod çözüldü.")
}

type xEncoder struct {
}

func (xencoder *xEncoder) Encode(value string) {
	fmt.Println("xEncoder ile kodlnadı.")
}
func (xencoder *xEncoder) Decode(value string) {
	fmt.Println("xEncoder ile kod çözüldü.")
}

/*
func main() {
	var encoder IEncoder
	encoder = &xEncoder{}

	encoder.Encode("123456")
	encoder.Decode("qweqwe1ğpo2qeqw")

	/*
		var encoder *xEncoder
		encoder = &xEncoder{}
		encoder.Encode("123456")
		encoder.Decode("xasdsawqe1")
}
*/
