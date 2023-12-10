package main

import "fmt"

var productName [4] string
var price [4] float32

func main() {
	productName[0] = "Macbook"
	productName[1] = "Asus"
	productName[2] = "Acer"
	productName[3] = "Hp"

	price := [4]float32{4000,2000,1500,6000} 

	fmt.Println(productName)
	fmt.Println(price)
}