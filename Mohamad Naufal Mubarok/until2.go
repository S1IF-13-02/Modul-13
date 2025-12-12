package main

import (
	"fmt"
	"math"
)

func main() {
	var bilangan float64

	fmt.Println("Masukkan bilangan desimal:")
	fmt.Scan(&bilangan)
	
	target := math.Ceil(bilangan)
	
	var total float64 = 0
	var step float64 = 0.1 
	
	fmt.Printf("\nMasukan: %.1f\n", bilangan)
	fmt.Printf("Target (pembulatan keatas): %.0f\n\n", target)
	fmt.Println("Keluaran:")
	
	for {
		total += step
		
		if total < target {
			fmt.Printf("%.1f\n", total)
		} else {
			fmt.Printf("%.0f\n", target)
			break
		}
	}
	
	fmt.Printf("\nProses selesai! Target %.0f tercapai.\n", target)
}