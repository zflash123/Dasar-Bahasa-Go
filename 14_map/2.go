package main

import "fmt"

func main() {
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)

	fmt.Print(chicken3["a"])
	fmt.Print(chicken4["a"])
	fmt.Print(chicken5["a"])
}
