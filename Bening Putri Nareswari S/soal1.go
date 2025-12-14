package main

import (
	"fmt"
)

func main() {
	var number int
	var digit int = 0

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)

	if number == 0 {
		digit = 1
	}

	for number > 0 {
		number = number / 10

		digit++
	}

	fmt.Printf("Banyaknya digit adalah: %d\n", digit)
}
