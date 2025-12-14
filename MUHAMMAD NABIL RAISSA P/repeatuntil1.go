package main
import "fmt"

func main() {
    var kata string
    var jumlah int

    fmt.Print("Masukan kata : ")
    fmt.Scan(&kata)

    fmt.Print("Jumlah : ")
    fmt.Scan(&jumlah)

    for i := 0; i < jumlah; i++ {
		fmt.Println(kata)
	}
}