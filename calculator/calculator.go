package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

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

func getOperator() string {
	fmt.Print(" + - * / ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add (val1, val2 float64) float64 {
	return val1 + val2
}
func minus (val1, val2 float64) float64 {
	return val1 - val2
}
func multiple (val1, val2 float64) float64 {
	return val1 * val2
}
func divide (val1, val2 float64) float64 {
	return val1 / val2
}

func main() {
	value1 := getInput("value1 =")
	value2 := getInput("value2 =")
	
	var result float64 
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
	fmt.Printf("result is %v", result)
}
