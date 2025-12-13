package main

import "fmt"

func main(){
	var x, y int
	var kondisi bool

	fmt.Scan(&x, &y)

	for {
		x -= y
		fmt.Println(x)
		if x == 0{
			kondisi = true
			break
		}
		if x < 0 {
			kondisi = false
			break
		}
	}
	fmt.Println(kondisi)
}