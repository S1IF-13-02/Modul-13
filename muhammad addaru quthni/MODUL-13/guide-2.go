package main 

import "fmt"

func main(){

	var bilangan int 
	var continueLoop bool

	
	for continueLoop = true; continueLoop;{
		fmt.Print("Masukkan bilangan bulat positif: ")
		fmt.Scan(&bilangan)
		continueLoop = bilangan <= 0

		if bilangan <= 0{
			fmt.Printf("%d bukan bilangan bulat positif. Silahkan coba lagi\n", bilangan)
		}
	}
	fmt.Printf("%d adalah bilangan bulat positif", bilangan)
}