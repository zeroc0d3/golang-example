// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

func main() {
	fmt.Println("Print Alphabet Only")
	fmt.Println("=========================")
	loopAlphabet()
	fmt.Println("\n")
	fmt.Println("Print Number Only")
	fmt.Println("=========================")
	loopInteger()
	fmt.Println("\n")
	fmt.Println("Print Alphabet & Number")
	fmt.Println("=========================")
	loopAlphabetNumber()
}

func loopAlphabet() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf(" %c", i)
	}
}

func loopInteger() {
	for j := 0; j <= 9; j++ {
		fmt.Printf(" %d", j)
	}
}

func loopAlphabetNumber() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf(" %c", i)
		if i == 'Z' {
			loopInteger()
		}
	}
}
