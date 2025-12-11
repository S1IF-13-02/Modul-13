package main

import "fmt"

func main() {
	var n string
	fmt.Println("Masukkan Bilangan:")
	fmt.Scan(&n)
	fmt.Println("Keluaran:")
	fmt.Println(len(n))
}