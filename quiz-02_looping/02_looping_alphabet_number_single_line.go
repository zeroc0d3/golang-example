// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

func main() {
	loopAlphabet()
	//fmt.Println("\n")
	//loopInteger()
}

func loopAlphabet() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf(" %c", i)
		if i == 'Z' {
			for j := 0; j <= 9; j++ {
				fmt.Printf(" %d", j)
			}
		}
	}
}

func loopInteger() {
	for i := 0; i <= 9; i++ {
		fmt.Printf(" %d", i)
	}
}
