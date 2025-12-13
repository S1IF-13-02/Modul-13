package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	hitung:= 0
	for{
		n = n/10
		hitung++

		if n == 0{
			break
		}
	}
	fmt.Println(hitung)
}