package main
import "fmt"
func main() {
	var kata string
	var jumlah int
	fmt.Scan(&kata, &jumlah)
	counter := 0
	for done := false; !done; {
		fmt.Println(kata)
		counter++
		done = (counter >= jumlah)
	}
}