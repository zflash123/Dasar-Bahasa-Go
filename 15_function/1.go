package main

import "fmt"

func main(){
	var names= []string{
		"Zaed",
		"Abdullah",
	}
	printMessage("halo", names)
}

func printMessage(message string, name []string ){
	var showName = name[0]+" "+name[1]
	fmt.Println(message, showName)
}