package main

import "fmt"

func main() {
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	var fourth, fifth, sixth = "empat", "lima", "enam"

	fmt.Println("This is a string without variables...")
	fmt.Println("Variable => first, second, third")
	fmt.Printf("Variable => %s, %s, %s \n", first, second, third)
	fmt.Println("---------------------------------------------")
	fmt.Println("Variable => fourth, fifth, sixth")
	fmt.Printf("Variable => %s, %s, %s", fourth, fifth, sixth)
}
