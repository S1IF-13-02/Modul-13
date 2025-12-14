package main

import "fmt"

func main() {
    var x, y int

	fmt.Print("Masukan x : ")
    fmt.Scan(&x)
	fmt.Print("Masukan y : ")
	fmt.Scan(&y)

    for {
        x = x - y
        fmt.Println(x)

        if x == 0 {
            fmt.Println(true) 
            break            
        }

        if x < 0 {
            fmt.Println(false) 
            break              
        }

    }
}
