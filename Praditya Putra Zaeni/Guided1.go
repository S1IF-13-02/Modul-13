package main

import "fmt"

func main() {
	var x string
	var y int
	fmt.Scan(&x, &y)
	
	i:=0
	for {
		fmt.Println(x)
		i++
		if!(i<y){
			break
		}
	}
}