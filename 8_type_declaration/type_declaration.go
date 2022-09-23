package main

import "fmt"

func main() {
	type NIM string
	var zaed NIM
	zaed = "21410391221"
	test := "Haloo"

	fmt.Printf("%T & %T", zaed, test) //main.NIM & string
}
