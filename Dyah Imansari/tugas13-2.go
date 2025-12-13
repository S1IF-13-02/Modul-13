package main
import (
	"fmt"
	"math"
)
func main() {
	var x float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&x)
	batas := math.Ceil(x)
	hasil := x
	fmt.Println("Hasil penjumlahan:")
	for {
		hasil += 0.1
		fmt.Printf("%.1f\n", hasil)
		if math.Round(hasil*10)/10 >= batas {
			break
		}
	}
	fmt.Printf("Hasil Pembulatan ke atas: %.0f\n", hasil)
}