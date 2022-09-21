package main

import "fmt"

func main(){
	const name = "Zaed"
	aChar := name[2] //will output ASCII "e" as decimal = 101
	fmt.Println(aChar) //101
	aString := string(aChar)
	fmt.Println(aString) //e
}