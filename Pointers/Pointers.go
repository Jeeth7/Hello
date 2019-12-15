package main

import "fmt"

func main() {
	var a int
	var b *int

	a = 10
	b = &a
	a = 12

	var j [6]int = [...]int{1, 2, 3, 4, 5, 6}
	var k *[6]int = &j
	fmt.Println(a, *b)
	fmt.Printf("\n%v %T", j, j)
	fmt.Printf("\n%v %T", *k, *k)
}
