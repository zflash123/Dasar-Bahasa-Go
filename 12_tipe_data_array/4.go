package main

import "fmt"

func main(){
	var nilai = []int{
		91,
		95,
		82,
	}

	fmt.Printf("nilaiku adalah %d, %d, %d \n", nilai[0], nilai[1], nilai[2])
	fmt.Printf("%T", nilai)
}