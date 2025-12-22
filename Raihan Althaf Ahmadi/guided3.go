package main

import "fmt"

func main() {
    var x, y int
    fmt.Print("Masukan Bilangan Positif (x y): ")
    fmt.Scan(&x, &y)

	a := (x % y == 0)
    for {
        x -= y
        fmt.Println(x)
        if x % y == 0 {
			fmt.Println(a)
			break
        }
        if x < 0 {
			fmt.Println(a)
            break
        }
    }
}
