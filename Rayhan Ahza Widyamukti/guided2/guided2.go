package main

import (
	"fmt"
)

func main() {
	var n int

	for {
		fmt.Print("Masukkan Bilangan: ")
		fmt.Scan(&n)

		if n > 0 {
			break
		}

	}
	fmt.Println(n, "merupakan bilangan bulat positif")

}
