package main

import "fmt"

var product = make(map[string]float64)

func main() {
	//add
	product["Macbook"] = 40000
	product["Iphone"] = 20000
	product["Ipad"] = 15000
	fmt.Println("product =", product)

	//dalete
	delete(product, "Ipad")
	fmt.Println("product =", product)

	//update
	product["Iphone"] = 36000
	fmt.Println("product =", product)

	value1 := product["Iphone"]
	fmt.Println("Price of iphone :", value1)

	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println(courseName)
}