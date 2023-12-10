package main

import "fmt"
func hello() {
	fmt.Println("Hello New")
}

func plus(value1, value2 float32) float32{
	// result := value1 + value2
	// fmt.Println("result =",result)

	return value1 + value2
}

func minus(value1, value2, value3 int) int {
	return value1 - value2 - value3
}

func main() {
	hello()

	fmt.Println(plus(2.1,3.5))

	resultMinus := minus(6,3,2)
	fmt.Println("result minus :" ,resultMinus)
}