package main

import "fmt"

func main() {
    var num float64
    fmt.Scan(&num)
    
    current := int(num * 10)
    target := ((current / 10) + 1) * 10
    
    first := true
    for done := false; !done; {
        current += 1
        
        if !first {
            fmt.Print(" ")
        }
        
        if current % 10 == 0 {
            fmt.Print(current / 10)
        } else {
            fmt.Printf("%d.%d", current/10, current%10)
        }
        
        first = false
        done = current >= target
    }
}