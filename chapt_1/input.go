package main

import "fmt"

func UserInput() {
	fmt.Printf("Please enter your name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)
}
