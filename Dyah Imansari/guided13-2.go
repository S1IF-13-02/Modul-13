package main
import "fmt"
func main() {
	var angka int
	var done bool
	for done = false; !done; {
		fmt.Scan(&angka)
		done = (angka > 0)
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", angka)
}