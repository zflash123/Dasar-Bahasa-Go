package main

import "fmt"

func main(){
	var fruits [7]string
	fruits[0] = "Grape"
	fruits[1] = "Orange"
	fruits[2] = "Melon"
	fruits[3] = "Jackfruit"
	fruits[4] = "Watermelon"
	fruits[5] = "Apple"
	fruits[6] = "Duku"

	var sliceFruit = fruits[1:4]
	sliceFruit2 := append(sliceFruit, "Manggo")
	
	fmt.Println(sliceFruit2)
	fmt.Println("Length of sliceFruit:", len(sliceFruit))
	fmt.Println("Capasity of sliceFruit:", cap(sliceFruit))
}
