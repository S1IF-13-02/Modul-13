package main

import "fmt"

func main(){
	var x string
	var y int
	var piw bool

	fmt.Scan(&x, &y)

	for piw = false; !piw; {
		fmt.Println(x)
		y--
		piw = y < 1 
	}
}