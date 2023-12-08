package main

import (
	"fmt"
	"strings"
)

func add(a, b int) int {
	result := a + b
	return result
}

// variadic function
func join(element ...string) string {
	return strings.Join(element, "")
}

func main() {
	result := add(1, 2)
	fmt.Println(result)

	fmt.Println(join("Manh", "Pham"))

	// anonymous function
	func() {
		fmt.Println("I am a anonymous function")
	}()

	// anonymous function with arguments
	func(ele string) {
		result := fmt.Sprintf("I am a anonymous function with %s", ele)
		fmt.Println(result)
	}("Manh")
}
