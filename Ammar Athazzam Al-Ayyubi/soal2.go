package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&a)

	batas := math.Ceil(a)
	for {
		a = a + 0.1
		if a < batas {
			fmt.Printf("%.1f\n", a)
		} else {
			fmt.Println(int(batas))
			break
		}
	}
}
