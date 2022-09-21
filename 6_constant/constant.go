package main

import "fmt"

func main(){
	const firstName = "Zaed"
	const lastName = "Abdullah"

	fmt.Println(firstName)
	//error karena const nilainya tidak bisa dirubah setelah diinisialisasi
	firstName = "Eko"
}