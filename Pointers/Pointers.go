package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var a int
	var b *int

	a = 10
	b = &a
	a = 12

	var j [6]int = [...]int{1, 2, 3, 4, 5, 6}
	var k *[6]int = &j
	var l interface{} = k
	p1 := person{
		name: "Inderjit",
		age:  34,
	}
	fmt.Println(a, *b)
	fmt.Printf("\n%v %T", j, j)
	fmt.Printf("\n%v %T", *k, *k)
	fmt.Printf("\n%v %T", l, l)
	fmt.Printf("\n")
	fmt.Println(p1)
}
