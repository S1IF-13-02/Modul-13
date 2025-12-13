package main

import "fmt"

func main() {

	var x, y int
	fmt.Println("Masukan nilai x dan y:")
	fmt.Scan(&x, &y)

	for {
		fmt.Println("Nilai x setelah dikurangi y:")
		x = x - y
		fmt.Println(x)

		if x <= 0 {
			break
		}
	}

	fmt.Println(x == 0)

}
