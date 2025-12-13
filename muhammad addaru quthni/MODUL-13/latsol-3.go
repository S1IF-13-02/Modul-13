package main

import "fmt"

func main() {
    var target, donation, total, donorCount int
    
    fmt.Scan(&target)
    
    total = 0
    donorCount = 0
    
    for done := false; !done; {
        fmt.Scan(&donation)
        donorCount++
        total += donation
        
        fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donorCount, donation, total)
        
        done = total >= target
    }
    
    fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, donorCount)
}