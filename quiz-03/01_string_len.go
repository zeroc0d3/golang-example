package main

import "fmt"
import "unicode/utf8"

func main() {
	fmt.Println("Hello, 世界", len("世界"), utf8.RuneCountInString("世界"))
	fmt.Println(len([]rune("世界")))
}
