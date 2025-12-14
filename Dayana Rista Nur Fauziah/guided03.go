package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("masukan nilai x : ")
	fmt.Scan(&x)
	fmt.Print("masukan nilai y : ")
	fmt.Scan(&y)

	for x > 0 {
		x = x - y //untk menghitung
		fmt.Println(x)
	}

	if x == 0 { // kalau x hasilnya positif maka akan mencetak true
		fmt.Println("true") // kalau negatif maka akan mencetak negatif
	} else {
		fmt.Println("false")
	}
}
