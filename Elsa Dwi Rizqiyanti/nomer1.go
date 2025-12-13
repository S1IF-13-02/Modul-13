package main

import "fmt"

func main() {
	var n int

	for {
		fmt.Print("Masukkan bilangan bulat positif: ")
		fmt.Scan(&n)

		if n > 0 {
			break
		}

		fmt.Println("Input harus bilangan positif!")
	}

	count := 0
	tmp := n

	for tmp > 0 {
		tmp /= 10
		count++
	}

	fmt.Printf("Banyaknya digit: %d\n", count)
}
