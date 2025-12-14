package main

import "fmt"

func main() {
	var kata string
	var n int

	fmt.Print("masukan kata : ")
	fmt.Scan(&kata)
	fmt.Print("masukan jumlah kata : ")
	fmt.Scan(&n)
	i := 0

	for {
		fmt.Println(kata)
		i++

		if i == n {
			break
		}

	}

}
