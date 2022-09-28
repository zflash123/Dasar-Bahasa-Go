package main

import "fmt"

func main(){
	var nilai = [...]int{
		91,
		95,
		82,
		83,
	}
	fmt.Println(nilai)
	fmt.Println(len(nilai))
	nilai[0] = 99
	fmt.Printf("nilaiku adalah %d, %d, %d", nilai[0], nilai[1], nilai[2])
}