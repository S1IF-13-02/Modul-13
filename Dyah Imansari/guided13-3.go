package main
import "fmt"
func main() {
	var X, Y int
	var done bool
	fmt.Print("Masukkan bilangan X dan Y: ")
	fmt.Scan(&X, &Y)
	fmt.Println("Proses pengurangan: ")
	for done = false; !done; {
		X = X - Y
		fmt.Println(X)
		done = (X <= 0)
	}
	fmt.Println("Apakah X kelipatan dari Y?")
	fmt.Println(X == 0)
}