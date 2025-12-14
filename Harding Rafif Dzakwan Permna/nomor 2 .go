package main // program utama Go

import "fmt" // untuk input/output

func main() {
    var x float64
    fmt.Print("Masukkan bilangan desimal: ")
    fmt.Scan(&x)    // baca input, misalnya 2.7

    // --- hitung batas atas (pembulatan ke atas) ---
    batas := int(x)             // buang desimal: 2.7 -> 2
    if float64(batas) < x {     // kalau 2 < 2.7
        batas = batas + 1       // jadikan 3 (ceil)
    }

    nilai := x                  // mulai dari nilai awal: 2.7

    for nilai < float64(batas) { // selama nilai masih < 3
        nilai = nilai + 0.1      // naikkan 0.1: 2.8, 2.9, 3.0, ...
        
        if nilai < float64(batas) {     
            // selama belum sampai batas (masih 2.8, 2.9) -> cetak 1 angka di belakang koma
            fmt.Printf("%.1f\n", nilai)  // contoh: 2.8, 2.9
        } else {
            // kalau sudah mencapai batas (3) -> cetak tanpa desimal
            fmt.Printf("%.0f\n", nilai)  // cetak: 3
        }
    }
}
