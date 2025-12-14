package main

import "fmt"

func main(){
	var x, y int
	fmt.Scan(&x, &y)
	done := false
	for !done {
		x = x - y 
		fmt.Println(x)
		done = x <= 0  
	}
	fmt.Println(x == 0)
}