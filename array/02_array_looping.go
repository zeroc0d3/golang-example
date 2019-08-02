package main

import "fmt"

func main() {
	loopingArray()
	fmt.Println("----------------------------")
	loopingArrayAgain()
}

func loopingArray() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d: %s \n", i, fruit)
	}
}

func loopingArrayAgain() {
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
