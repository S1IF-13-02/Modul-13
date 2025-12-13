package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)

	for x > 0 {
		fmt.Println(x)
		x = x - y

	}
	fmt.Println(x)

	if x == 0 {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}
