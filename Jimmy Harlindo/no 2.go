package main

import "fmt" 


func main() {
	var x float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&x)

	batas := math.Ceil(x) 
	curr := x

	for {
		curr += 0.1
		fmt.Printf("%.1f\n", curr)

		if curr >= batas {
			break
		}
	}
}
