package main

import "fmt"

// 3.Array
var productName [4]string

func main() {
	productName[0] = "Macbook"
	productName[1] = "iPad"
	productName[2] = "iPhone"
	productName[3] = "AirPods"

	// ประกาศตัวแปรของ array แบบ short form
	price := [4]float32{40000, 30000, 20000, 2000}

	fmt.Println(productName)
	fmt.Println(price)
}
