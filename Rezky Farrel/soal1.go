package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	count := 0
	temp := n

	for temp > 0 {
		temp /= 10
		count++
	}

	fmt.Println("Jumlah digit:", count)
}
