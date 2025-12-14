package main

import "fmt"

func main() {
	var n int

	for {
		fmt.Print("Masukkan bilangan: ")
		fmt.Scan(&n)
		if n < 0 {
			continue
		}
		if n > 0 {
			break
		}
	}
	fmt.Println(n, "adalah bilangan bulat positif")
}
