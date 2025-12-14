package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)

	i := 0
	for {
		n /= 10
		i++

		if n == 0 {
			break
		}
	}

	fmt.Println(i)
}
