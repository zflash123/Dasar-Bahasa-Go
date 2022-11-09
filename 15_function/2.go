package main

import "fmt"

func main() {
	var names = []string{
		"Zaed",
		"Abdullah",
	}
	fmt.Println(printMessage("halo", names))
}

func printMessage(message string, name []string) string {
	var showName = name[0] + " " + name[1]
	return showName
}