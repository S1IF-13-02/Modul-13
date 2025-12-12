package main 

import "fmt"

func main(){

	var kata string
	var banyak_kata int

	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)
	fmt.Print("Masukkan banyak perulangan: ")
	fmt.Scan(&banyak_kata)

	i := 0

	for{
		fmt.Println(kata)
		i++

		if i == banyak_kata{
			break
		}
	}
}


