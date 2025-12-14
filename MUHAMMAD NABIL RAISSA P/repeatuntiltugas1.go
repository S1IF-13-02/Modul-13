package main
import "fmt"

func main(){
	var bilangan int
	fmt.Print(" Masukan Bilangan bulat : ")
	fmt.Scan(&bilangan)
	count := 0

	for {
		bilangan = bilangan/10
		count++ 

		if bilangan == 0{
			break
		}		
	}
	fmt.Print(count)
}