package main

import "fmt"

func main() {
	var desimal float64
	var batas, hitung int
	fmt.Scan(&desimal)

	hitung = int(desimal * 10)
	batas = (hitung / 10) * 10
	if hitung%10 != 0 {
		batas += 10
	}

	hitung += 1
	for {
		fmt.Printf("%.1f\n", float64(hitung)/10)
		if hitung >= batas {
			break
		}
		hitung++
	}
}
