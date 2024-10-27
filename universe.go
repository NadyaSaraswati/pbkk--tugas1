package main

import (
	"fmt"
)

func main() {
	var userInput int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&userInput)

	if userInput == 42 {
		fmt.Println("Hello, Universe")
	} else {
		fmt.Println(userInput)
	}
}