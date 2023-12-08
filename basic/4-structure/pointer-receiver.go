package main

import "fmt"

type Vector struct {
	x int
	y int
}

func (p *Vector) Desc() {
	fmt.Println(p.x)
	fmt.Println(p.y)
}

func main() {
	x := &Vector{1, 2}
	x.Desc()
}
