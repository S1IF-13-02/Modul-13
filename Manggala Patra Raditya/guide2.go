package main

import (
	"fmt"
)

func main() {
	var n int

	for {
		fmt.Scan(&n)
		if n > 0 {
			break
		}
	}

	fmt.Printf("%d adalah bilangan bulat positif\n", n)
}