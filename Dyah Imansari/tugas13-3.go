package main
import "fmt"
func main() {
	var target, donasi int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scanln(&target)
	fmt.Println("Masukkan donasi dari tiap donatur: ")
	counter := 0
	total := 0
	for {
		counter++
		fmt.Scan(&donasi)
		fmt.Printf("Donatur %d: Menyumbang %d. ", counter, donasi)
		total += donasi
		fmt.Printf("Total terkumpul: %d\n", total)
		if total >= target {
			break
		}
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.", total, counter)
}