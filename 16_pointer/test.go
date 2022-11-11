package main

import "fmt"

func main(){
	var nilai1 = 90
	var pNilai1 = &nilai1
	fmt.Printf("pointer pNilai: %p, dengan tipe-data: %T \n", pNilai1, pNilai1)

	nilai1 = 68
	fmt.Printf("pointer pNilai: %p, dengan tipe-data: %T, dengan value: %v\n", pNilai1, pNilai1, *pNilai1)
}