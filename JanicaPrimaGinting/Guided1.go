package main

import "fmt"

func main(){
	var kata string
	var jmlh int
	fmt.Scan(&kata)
	fmt.Scan(&jmlh)
	done := false
	for !done {
		fmt.Println(kata)
		jmlh = jmlh - 1
		done = jmlh < 1
	}
}