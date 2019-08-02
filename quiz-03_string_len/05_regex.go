package main

import (
	"regexp"
	"strings"
	"unicode"
)

func main() {

	str := "\u0308" + "a\u0308" + "o\u0308" + "u\u0308"
	str2 := "a" + strings.Repeat("\u0308", 1000)

	println(4 == GraphemeCountInString(str))
	println(4 == GraphemeCountInString2(str))

	println(1 == GraphemeCountInString(str2))
	println(1 == GraphemeCountInString2(str2))

	println(true == IsStreamSafeString(str))
	println(false == IsStreamSafeString(str2))
}

func GraphemeCountInString(str string) int {
	re := regexp.MustCompile("\\PM\\pM*|.")
	return len(re.FindAllString(str, -1))
}

func GraphemeCountInString2(str string) int {

	length := 0
	checked := false
	index := 0

	for _, c := range str {

		if !unicode.Is(unicode.M, c) {
			length++

			if checked == false {
				checked = true
			}

		} else if checked == false {
			length++
		}

		index++
	}

	return length
}

func IsStreamSafeString(str string) bool {
	re := regexp.MustCompile("\\PM\\pM{30,}")
	return !re.MatchString(str)
}
