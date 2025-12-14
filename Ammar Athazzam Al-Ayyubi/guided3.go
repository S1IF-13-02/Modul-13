package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	for {
		a = a - b
		fmt.Println(a)
		if a-b == 0 {
			fmt.Println(true)
			break
		
		} else if a < 0 {
			fmt.Println(false)
			break
		}
	}
}
