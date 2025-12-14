package main

import "fmt"

func main() {
	var bilangan int

	for {
		fmt.Scan(&bilangan) //-2,-1,0,1
		if bilangan > 0 {
			break //untuk memberhentikan
		}
	}

	fmt.Println(bilangan, "adalah bilangan bulat positif.")
}
