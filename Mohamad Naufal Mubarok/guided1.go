package main

import "fmt"

func main() {
	var kata string
	var n, i int

	fmt.Print("Masukan kata : ")
	fmt.Scan(&kata)
	fmt.Print("Masukan pengulangan : ")
	fmt.Scan(&n)

	i = 0
	for {
		fmt.Println(kata)
		i++
		if i >= n {
			break			
		}
	}

}