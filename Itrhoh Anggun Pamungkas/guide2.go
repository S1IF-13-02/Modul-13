package main

import "fmt"

func main() {
	var bil int

	for {
		fmt.Scan(&bil)
		
		if bil > 0 {
			break
		}
	}
	fmt.Println(bil, "adalah bilangan positif")
}
