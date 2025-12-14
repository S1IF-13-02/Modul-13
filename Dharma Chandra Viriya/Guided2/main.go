package main

import "fmt"

func main() {
	var n int

	for {
		fmt.Print("Masukkan angka: ")
		fmt.Scanln(&n)
		if n > 0 {
			fmt.Printf("%d adalah bilangan bulat positif\n", n)
			break
		}
	}
}
