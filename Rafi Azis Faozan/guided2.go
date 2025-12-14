package main

import "fmt"

func main() {
	var bilangan int
	for {
		fmt.Print("Masukkan bilangan bulat positif: ")
		fmt.Scan(&bilangan)
		if bilangan > 0 {
			break
		}
	}
	fmt.Print(bilangan, " adalah bilangan bulat positif")

}
