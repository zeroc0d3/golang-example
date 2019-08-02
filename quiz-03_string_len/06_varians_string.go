package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//text_sample := "这是个测试"
	text_sample := "世界"
	len1 := len([]rune(text_sample))
	len2 := bytes.Count([]byte(text_sample), nil) - 1
	len3 := strings.Count(text_sample, "") - 1
	len4 := utf8.RuneCountInString(text_sample)

	fmt.Println("Get string length with various way: ")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Method                                      | Result        ")
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("len([]rune(text_sample))                    | %d \n", len1)
	fmt.Printf("bytes.Count([]byte(text_sample), nil) - 1   | %d \n", len2)
	fmt.Printf("strings.Count(text_sample, '') - 1          | %d \n", len3)
	fmt.Printf("utf8.RuneCountInString(text_sample)         | %d \n", len4)
	fmt.Println("------------------------------------------------------------")
}
