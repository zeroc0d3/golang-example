package main

import "fmt"

func main() {
	var firstName string = "Johny"
	lastName := "Wick" //---> type inference without "var"
	fmt.Printf("This is a string without variables ... \n")
	fmt.Printf("Welcome, %s %s", firstName, lastName)
}
