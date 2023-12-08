package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(2) == 1
	fmt.Println(random)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	switch day := 1; day {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}
}
