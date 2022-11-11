package main

import "fmt"

func main(){
	type bicycle struct{
		name string;
		front_gear int;
		back_gear int;
	}
	var b1 bicycle

	b1.name = "Trek Emonda"
	b1.front_gear = 3
	b1.back_gear = 11
	fmt.Printf("Halo namaku Kiyotaka \nSepedaku: %v \n", b1.name)
	fmt.Printf("Speed sepedaku: %vx%v", b1.front_gear, b1.back_gear)
}