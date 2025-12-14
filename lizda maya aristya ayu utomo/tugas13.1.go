package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0

	for {
		n = n / 10
		count++
		if n == 0 {
			break
		}
	}
	fmt.Println(count)
}
