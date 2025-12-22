package main

import "fmt"
func main (){
	var x int
	for {
		fmt.Print("masukan angka : ")
		fmt.Scan(&x)
		if x > 0 {
			fmt.Print(x," adalah bilangan bulat positif")
			break	
		}
	}
}