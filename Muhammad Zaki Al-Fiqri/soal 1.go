package main

import "fmt"

func main() {
	var bilanganPositif int
	var jumlahDigit int = 0

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&bilanganPositif)

	if bilanganPositif == 0 {
		jumlahDigit = 1
	} else {
		for {
			bilanganPositif = bilanganPositif / 10
			jumlahDigit++

			if bilanganPositif <= 0 {
				break
			}
		}
	}

	fmt.Printf("%d", jumlahDigit)
}
