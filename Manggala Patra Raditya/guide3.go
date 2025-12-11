package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Println("Keluaran:")

	for x > 0 {
		x -= y
		fmt.Println(x)
	}

	if x == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}