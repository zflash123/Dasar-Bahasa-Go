package main

import "fmt"

func main(){
	type bicycle struct{
		name string;
		front_gear int;
		back_gear int;
	}
	var b1 = bicycle{
		"Polygon",
		3,
		10,
	}
	fmt.Println(b1.name)

	var b2 = bicycle{
		front_gear: 3,
	}
	fmt.Println(b2.front_gear)

	var b3 = bicycle{}
	b3.name = "Pinarello"
	fmt.Println(b3.name)
}