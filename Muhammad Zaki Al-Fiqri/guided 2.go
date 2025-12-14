package main

import "fmt"

func main() {

	var n int

	for {
		fmt.Print("masukan bilangan bulat positif : ")
		fmt.Scan(&n)

		if n > 0 {
			break
		}
		fmt.Println("bukan bilangan bulat positif")
	}
	fmt.Printf("%d adalah bilangan bulat positif", n)
}
