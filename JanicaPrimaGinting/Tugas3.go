package main

import "fmt"

func main(){
	var tDonasi, donasi int
	fmt.Scan(&tDonasi)
	donatur := 0
	tTerkumpul := 0
	tTercapai := false
	for !tTercapai {
		fmt.Scan(&donasi)
		donatur = donatur + 1
		tTerkumpul = tTerkumpul + donasi
		fmt.Println("Donatur", donatur,": Menyumbang", donasi,". Total terkumpul:", tTerkumpul)
		tTercapai = tTerkumpul >= tDonasi
	}
	fmt.Println("Target tercapai! Total donasi:", tTerkumpul, "dari", donatur, "donatur.")
}