package main

import "fmt"

func main() {
    var x, y int
	fmt.Print("Masukan niai x : ")
	fmt.Scan(&x)
	fmt.Print("Masukan nilai y :")
	fmt.Scan(&y)

    for x > 0 {
        x = x - y
		fmt.Println(x)
    }

    if x == 0 {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}