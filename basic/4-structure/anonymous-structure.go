package main

import "fmt"

func main() {
	persion := struct {
		name string
		age  int
	}{
		name: "Manh",
		age:  30,
	}

	fmt.Println(persion)
}
