package main

import "fmt"
import "github.com/rivo/uniseg"

func main() {
	gr := uniseg.NewGraphemes("👍🏼!")
	for gr.Next() {
		fmt.Printf("%x ", gr.Runes())
	}
	// Output: [1f44d 1f3fc] [21]
}
