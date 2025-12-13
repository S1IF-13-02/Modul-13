package main

import "fmt"

func main() {
	var bilangan, hitung, temp int
	fmt.Scan(&bilangan)

	hitung = 0
	temp = bilangan

	for {
		temp = temp / 10
		hitung++
		if temp == 0 {
			break
		}
	}
	fmt.Println(hitung)
}
