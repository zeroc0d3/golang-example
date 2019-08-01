package main

import "fmt"

/* Print Alphabet */
func printAlphabet() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf(" %c", i)
	}
}

/* Print Numbers */
func printNumbers() {
	for i := 0; i <= 9; i++ {
		fmt.Printf(" %d", i)
	}
}

/* Main Program */
func main() {
	fmt.Println("Looping print Alphabet")
	fmt.Println("==================================")
	printAlphabet()
	fmt.Println("\n")
	fmt.Println("Looping print Numbers")
	fmt.Println("==================================")
	printNumbers()
}
