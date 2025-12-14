package main

import "fmt"

func main() {
	var angka int
	var cuaca string

	fmt.Scan(&cuaca, &angka)

	i := 0
	for kondisi := false; !kondisi; {
		fmt.Println(cuaca)
		i++
		kondisi = i >= angka
	}
}