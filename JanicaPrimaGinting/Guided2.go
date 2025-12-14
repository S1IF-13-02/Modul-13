package main

import "fmt"

func main(){
	var n int
	for done := false ; !done ; {
		fmt.Scan(&n)
		done = n > 0
	}
	fmt.Println(n,"adalah bilangan bulat positif")
}