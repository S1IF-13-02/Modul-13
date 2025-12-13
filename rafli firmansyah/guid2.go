package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for {
		if n > 0 {
			break
		}
	}
	fmt.Print(n, "adalah bilangan bulat positif")
}
