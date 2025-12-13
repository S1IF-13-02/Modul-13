package main

import "fmt"

func main() {
    var num, count int
    
    fmt.Scan(&num)
    
    if num == 0 {
        fmt.Println(1)
        return
    }
    
    count = 0
    for done := false; !done; {
        num = num / 10
        count++
        done = num == 0
    }
    
    fmt.Println(count)
}