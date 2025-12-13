package main
import "fmt"
func main() {
	var x int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&x)
	counter := 0
	for x != 0 {
		counter++
		x = x/10
	}
	fmt.Println("Jumlah digit:", counter)
}