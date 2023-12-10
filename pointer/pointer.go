package main

import "fmt"

func zeroValue(iValue int) {
	 iValue = 0
}

func zeroPointer(iPointer *int) {
	*iPointer = 0
}

func main() {
	i := 1 
	fmt.Println("i :", i)

	zeroValue(i)
	fmt.Println("i from zeroValue =", i)
	
	zeroPointer(&i) //& = การเข้าถึง pointer
	fmt.Println("i value from zeroPointer =", i)
	fmt.Println("i address from zeroPointer =", &i)
}