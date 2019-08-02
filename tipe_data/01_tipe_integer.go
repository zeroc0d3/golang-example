package main

import "fmt"

var uint8_test uint8 = 255
var uint16_test uint16 = 65535
var uint32_test uint32 = 4294967295
var uint64_test uint64 = 18446744073709551615

var uint_test1 uint16 = uint16_test
var uint_test2 uint32 = uint32_test

var byte_test uint8 = uint8_test
var int8_test int8 = 127
var int16_test int16 = 32767
var int32_test int32 = 2147483647
var int64_test int64 = 9223372036854775807

var int_test1 = int32_test
var int_test2 = int64_test
var rune_test = int32_test

func main() {
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Type Integer (Min - Max)   ")
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("uint8  (0 <-> %d) \n", uint8_test)
	fmt.Printf("uint16 (0 <-> %d) \n", uint16_test)
	fmt.Printf("uint32 (0 <-> %d) \n", uint32_test)
	fmt.Printf("uint64 (0 <-> %d) \n\n", uint64_test)
	fmt.Printf("uint   (0 <-> %d) \n", uint_test1)
	fmt.Printf("uint   (0 <-> %d) \n", uint_test2)
	fmt.Printf("byte   (0 <-> %d) \n\n", byte_test)
	fmt.Printf("int8   (-128 <-> %d) \n", int8_test)
	fmt.Printf("int16  (-32768 <-> %d) \n", int16_test)
	fmt.Printf("int32  (-2147483648 <-> %d) \n", int32_test)
	fmt.Printf("int64  (-9223372036854775808 <-> %d) \n\n", int64_test)
	fmt.Printf("int    (-2147483648 <-> %d) \n", int_test1)
	fmt.Printf("int    (-9223372036854775808 <-> %d) \n", int_test2)
	fmt.Printf("rune   (-2147483648 <-> %d) \n", rune_test)
	fmt.Println("-------------------------------------------------------")
}
