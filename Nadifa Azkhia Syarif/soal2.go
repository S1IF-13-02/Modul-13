package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	atas := int(x)
	if x > float64(atas) {
		atas = atas + 1
	}

	for {
		x = x + 0.1
		fmt.Printf("%.1f\n", x)

		if x >= float64(atas) {
			break
		}
	}
}
