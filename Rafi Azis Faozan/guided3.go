package main
import "fmt"
func main() {
	var x int
	var y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
    hasil := x
   	for {
        hasil -= y
        fmt.Println(hasil)

        if hasil <= 0 {  
            break
        }
    }

    fmt.Println(hasil == 0)
}