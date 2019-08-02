package main

import "fmt"

func main() {
	russian := "Спутник и погром"
	english := "Sputnik & pogrom"

	fmt.Println("count of bytes:",
		len(russian),
		len(english))

	fmt.Println("count of runes:",
		len([]rune(russian)),
		len([]rune(english)))
}
