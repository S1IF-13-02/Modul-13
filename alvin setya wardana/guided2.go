package main

import (
	"fmt"
)

func main() {
	var n int

	for {
		fmt.Print("Masukkan bilangan bulat positif: ")
		fmt.Scanln(&n)

		if n > 0 {
			break
		}
		fmt.Println("Input tidak valid. Silakan coba lagi.")
	}

	fmt.Printf("%d adalah bilangan positif\n", n)
}
