package main

import "fmt"
func main(){
	var nama string 
	var x int
	fmt.Print("Masukan Satu kata dan jumlah yang di ingin kan : ")
	fmt.Scan(&nama, &x)
	for {
		fmt.Println(nama)
		x--
		if x < 1{
			break
		}
	}
}