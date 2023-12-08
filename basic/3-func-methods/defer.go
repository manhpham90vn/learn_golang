package main

import "fmt"

func add(x, y int) {
	result := x + y
	fmt.Println(result)
}

func main() {
	add(1, 2)

	defer add(3, 4)

	add(5, 6)
}
