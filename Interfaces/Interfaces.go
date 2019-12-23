package main

import "fmt"

func main() {
	var test Writer = ConsoleWriter{
		name: "Inderjit",
		age:  33,
	}
	test.Write([]byte("Hello interface"))
	test.WriteObj()
	fmt.Println(test)
}

type Writer interface {
	Write([]byte) (int, error)
	WriteObj()
}

type ConsoleWriter struct {
	name string
	age  int16
}

func (feed ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func (feed ConsoleWriter) WriteObj() {
	feed.name = "New Name"
	fmt.Println(feed)
}

//func (feed *ConsoleWriter) WriteObjPointer() {
//	feed.name = "New Name"
//	fmt.Println(feed)
//}
