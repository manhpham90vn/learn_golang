package main

import (
	"basic/test"
	"fmt"
	"strconv"
)

// const
const PI = 3.14

func main() {
	fmt.Println("Hello World")
	fmt.Println(test.ExportedVariable)

	var localVariable = "Manh"
	fmt.Printf("%s\n", localVariable)

	var x uint8 = 20
	fmt.Println(x)

	var f float32 = 10.0
	fmt.Println(f)

	var isSuccess bool = true
	fmt.Println(isSuccess)

	// short variable
	name := "Manh"
	fmt.Println(name)

	a := int64(100)
	b := int64(200)

	// &
	and := a & b
	fmt.Printf("a = %s b = %s result = %s %d\n", strconv.FormatInt(a, 2), strconv.FormatInt(b, 2), strconv.FormatInt(and, 2), and)

	// or
	or := a | b
	fmt.Printf("a = %s b = %s result = %s %d\n", strconv.FormatInt(a, 2), strconv.FormatInt(b, 2), strconv.FormatInt(or, 2), or)

	// xor
	xor := a ^ b
	fmt.Printf("a = %s b = %s result = %s %d\n", strconv.FormatInt(a, 2), strconv.FormatInt(b, 2), strconv.FormatInt(xor, 2), xor)

	// <<
	left := a << 1
	fmt.Printf("left = %s %d\n", strconv.FormatInt(left, 2), left)

	// >>
	right := a >> 1
	fmt.Printf("right = %s %d\n", strconv.FormatInt(right, 2), right)
}
