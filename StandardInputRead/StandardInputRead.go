package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i int
	var j string
	fmt.Println("Enter a number and a string")
	n, e := fmt.Scanf("%d %s", &i, &j)
	if e != nil {
		fmt.Println("error")
	}
	fmt.Println(n, i, j)

	fmt.Println("Enter a number and a string")
	reader := bufio.NewReader(os.Stdin)
	i1, _ := reader.ReadString('\n')
	fmt.Println(i1)

}
