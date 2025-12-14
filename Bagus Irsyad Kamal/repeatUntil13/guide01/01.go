package main

import "fmt"

func main() {
    var word string 
    var repetitions int 
    fmt.Print("Masukan kata :")
    fmt.Scan(&word)
    fmt.Print("Masukan berapa kali pengulanggan :")
    fmt.Scan(&repetitions) 

    counter := 0 
    for done := false; !done; { 
        fmt.Println(word) 
        counter++ 
        done = (counter >= repetitions) 
    }
    fmt.Println("Selesai mencetak.")
}