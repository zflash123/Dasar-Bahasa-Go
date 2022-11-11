package main

import "fmt"

func main(){
	type laptop struct{
		brand string;
		ram int;
		cpu string;
	}
	var l1 = laptop{
		"Asus",
		4,
		"Core i5",
	}
	var pointerL1 *laptop = &l1
	fmt.Println(pointerL1)

}