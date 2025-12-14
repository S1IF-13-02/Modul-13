package main

import "fmt"

func main() {
    var bilangan float64
    fmt.Scan(&bilangan)
    
    uNilai := int(bilangan * 10)
    nDepan := uNilai / 10
    nAkhir := (nDepan + 1) * 10

    kondisi := false

    for !kondisi {
        uNilai = uNilai + 1
        if uNilai % 10 == 0 {
            fmt.Println(uNilai / 10)
        }else{
            fmt.Printf("%.1f\n",float64(uNilai)/10)
        }
        if uNilai >= nAkhir {
            kondisi = true
        }
    }
}