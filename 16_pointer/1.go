package main

import "fmt"

func main(){
	var nilai = 90
	var pNilai = &nilai
	fmt.Printf("pointer pNilai: %p, dengan tipe-data: %T", pNilai, pNilai)
}