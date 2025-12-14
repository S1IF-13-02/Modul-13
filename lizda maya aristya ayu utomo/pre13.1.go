package main

import "fmt"

func main() {
	var k string
	var n int

	fmt.Scan(&k, &n)

	counter := 0

	for done := false; !done; {
		fmt.Println(k)
		counter++
		done = (counter >= n)
	}
}
