package main

import "fmt"

func main(){
	var target_donasi int

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target_donasi)

	total_donasi := 0
	donatur := 0

	for{
		var donasi int
		
		donatur++
		
		fmt.Printf("Donatur ke-%d: Menyumbang ", donatur)
		fmt.Scan(&donasi)

		if donasi <= 0 {
		fmt.Println("Donasi yang dimasukkan harus lebih dari 0.")
		donatur--
		continue
		}

		total_donasi += donasi

		if total_donasi >= target_donasi{
			break
		}
	}
	fmt.Printf("Target donasi tercapai. Total donasi: %d dari %d donatur\n", total_donasi, donatur)
}
