package main

import "fmt"

var message string = "OK"
var message2 = `Test Multiline. 
-------------------------------
If you can see this, this 
message write in multiline
`

func main() {
	fmt.Printf("Tipe string: %s \n\n", message)
	fmt.Printf("Multistring: %s \n", message2)
}
