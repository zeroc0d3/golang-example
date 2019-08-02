package main

import "fmt"
import "golang.org/x/text/unicode/norm"

func main() {
	var ia norm.Iter
	//ia.InitString(norm.NFKD, "école")
	ia.InitString(norm.NFKD, "世界")
	nc := 0
	for !ia.Done() {
		nc = nc + 1
		ia.Next()
	}
	fmt.Printf("Number of chars: %d\n", nc)
}
