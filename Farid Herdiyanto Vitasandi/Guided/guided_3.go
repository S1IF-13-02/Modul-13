package main 

import "fmt"

func main(){

	var x, y int 
	var isKelipatan bool

	fmt.Print("Masukkan bilangan (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan (y): ")
	fmt.Scan(&y)

	for isKelipatan = false; !isKelipatan;{
		x -= y
		fmt.Println(x)

		if x == 0{
			isKelipatan = true
			break
		}
		if x < 0 {
			isKelipatan = false
			break
		}
	}
	fmt.Println(isKelipatan)
}
