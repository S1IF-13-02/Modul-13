package main

import "fmt"

func main() {
	var text string
	var counter int

	fmt.Print("Masukkan teks: ")
	fmt.Scanln(&text)
	fmt.Print("Masukkan banyak keluaran (angka): ")
	fmt.Scanln(&counter)

	i := 0
	for {
		fmt.Println(text)
		i++
		if i == counter {
			break
		}
	}
}
