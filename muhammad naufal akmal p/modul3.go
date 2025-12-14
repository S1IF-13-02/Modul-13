package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Masukkan X: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan Y: ")
	fmt.Scan(&y)

	temp := x 

	for {
		temp = temp - y
		fmt.Println(temp)

		if temp <= 0 {
			break
		}
	}
	if temp == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}