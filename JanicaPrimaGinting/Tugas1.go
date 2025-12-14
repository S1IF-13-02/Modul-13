package main 

import "fmt"

func main(){
	var bilangan int
	fmt.Scan(&bilangan)
	digit := 0
	habis := false
	for !habis{
		bilangan = bilangan / 10
		digit ++
		habis = bilangan <= 0
	}
	fmt.Println(digit)
}