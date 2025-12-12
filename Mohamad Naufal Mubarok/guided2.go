package main

import (
	"fmt"
)


func main() {
	var n int

	for{
		fmt.Print("Masukan bilangan : ")
		fmt.Scan(&n)

		if n > 0 {
			break
		}
		fmt.Println(" bukan bilangan positif, coba lagi")
	}
		fmt.Println(n,"ini bilangan bulat positif")
}