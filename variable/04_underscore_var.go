package main

import "fmt"

func main() {
	_ = "This var never use"
	var firstName string = "John"
	lastName := "Doe"

	fmt.Println("This is a string without variable")
	fmt.Printf("First Name: %s \n", firstName)
	fmt.Printf("Last Name: %s \n", lastName)

	string_test, _ := "OK", "Never use"
	fmt.Printf("Test: %s", string_test)
}
