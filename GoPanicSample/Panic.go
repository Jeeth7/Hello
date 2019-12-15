package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start !")
	Panicer()
	fmt.Println("End !")
}

func Panicer() {
	fmt.Println("Start panicing  !")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("Something Happened")
	fmt.Println("End panicing !")
}
