package main
import "fmt"

func main() {
	var target int
	var donasi int
	var total int
	var jumlahDonatur int
	var selesai bool

	fmt.Scan(&target)

	total = 0
	jumlahDonatur = 0

	for selesai = false; !selesai; {
		fmt.Scan(&donasi)
		jumlahDonatur++
		total = total + donasi

		fmt.Printf(
			"Donatur %d : Menyumbang %d. Total terkumpul: %d\n",
			jumlahDonatur, donasi, total,
		)

		selesai = total >= target
	}

	fmt.Printf(
		"Target tercapai! Total donasi: %d dari %d donatur.\n",
		total, jumlahDonatur,
	)
}
