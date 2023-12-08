package main

import "fmt"

type Persion struct {
	name string
	age  int
}

func (p Persion) desc() {
	fmt.Println("desc func")
	fmt.Println("name: " + p.name)
	fmt.Printf("age: %d", p.age)
}

func main() {
	var test Persion
	fmt.Println(test)

	var a = Persion{"Manh", 33}
	b := Persion{"Ngan", 33}

	fmt.Println(a)
	fmt.Println(b.name)
	a.desc()

	c := &Persion{"M", 33}
	// truy cập thông thường
	fmt.Println((*c).age)
	// cũng ok
	fmt.Println(c.age)
}
