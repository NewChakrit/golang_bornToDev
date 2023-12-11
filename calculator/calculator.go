package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Validate input must be number 
func getInput (prompt string) float64{
	fmt.Printf("%v", prompt)
	input, _:= reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message,_ := fmt.Scanf("%v must number only", prompt)
		panic(message)
	}

	return value
}

// Number of input
func getNumInput (prompt string,) float64{
	fmt.Printf("%v", prompt)
	input, _:= reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message,_ := fmt.Scanf("%v must number only", prompt)
		panic(message)
	}

	return value
}

// Operator
func getOperator() string {
	fmt.Print("Operator : ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

//  Plus function
func add (val1, val2 float64) float64 {
	return val1 + val2
}

//  Minus function
func minus (val1, val2 float64) float64 {
	return val1 - val2
}

//  Multiple function
func multiple (val1, val2 float64) float64 {
	return val1 * val2
}

//  Divide function
func divide (val1, val2 float64) float64 {
	return val1 / val2
}

// Get result
func getResult() float64 {
	numOfInput := getNumInput("How many numbers do you want to calculate?")

	var result float64 

	for i := 1.0 ; i < numOfInput; i++ {
		var value1 float64
		if i == 1 {
			value1 = getInput("Number : ")
		}
		value1 = i
		value2 := getInput("Number : ")
		
		// Operator
		switch operator := getOperator(); operator {
		case "+":
			result = add(value1, value2)
		case "-":
			result = minus(value1, value2)
		case "*":
			result = multiple(value1, value2)
		case "/":
			result = divide(value1, value2)
		default:
			panic("wrong operator")
		}
	
	}
	return result

}

func main() {

	result := getResult()
	
	fmt.Printf("result is %v", result)
	
	
}
