package main

import "fmt"

// การประกาศตัวแปร

// 1.Integer
var numberInt, numberInt2 int = 1000, 2000

// 2. String
var msg string = "Hello "

func main() {
	// การประกาศตัวแปรแบบ short format
	numberfloat := 25.4

	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	// Variable Operation (int+int)
	fmt.Println(numberInt + numberInt2)

	// Variable Operation (int+float)
	fmt.Println(float64(numberInt) + numberfloat)

	// Variable Operation (str+str)
	fmt.Println(msg + "world")

	// Variable Operation (str+int)
	fmt.Println("my money =", numberInt)
}
