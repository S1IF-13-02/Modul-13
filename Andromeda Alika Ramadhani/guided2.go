package main

import "fmt"

func main() {
	var bil int
	var stop bool

	bil = 0
	for stop = false; !stop; {
		fmt.Scan(&bil)
		stop = bil > 0
	}

	fmt.Println(bil, "adalah bilangan bulat positif")
}
