package main

import "fmt"

func main() {
	looping_array()
	fmt.Println("----------------------------")
	looping_array_again()
}

func looping_array() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d: %s \n", i, fruit)
	}
}

func looping_array_again() {
	/* horizontal */
	var fruits = [4]string{"orange", "durian", "mango", "papaya"}

	/* vertical */
	/*
		var fruits = [4]string{
			"orange",
			"durian",
			"mango",
			"papaya"
		}
	*/

	for _, fruit := range fruits {
		fmt.Printf("name fruit: %s \n", fruit)
	}
}
