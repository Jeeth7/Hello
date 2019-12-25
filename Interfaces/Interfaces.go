package main

import (
	"bytes"
	"fmt"
)

func main() {
	var test Writer = ConsoleWriter{
		name: "Inderjit",
		age:  33,
	}
	test.Write([]byte("Hello interface"))

	var test1 Printer
	test1, ok := test.(Printer) //Type conversion
	if ok {
		fmt.Println("conversion success")
	}
	test1.Print()
	fmt.Printf("\nPrint type converson of Writer to Printer \nWriter %v\n Printer %v", test, test1)

	jj := test.(ConsoleWriter)
	var test2 WritePointer = &jj
	test2.WriteObjPointer("Inderjit Kolappan")
	fmt.Printf("\nPrint type converson of Writer to Printer \nWritePointer %v\n ConsoleWriter %v", test2, jj)

	fmt.Printf("\nConsoleWriter jj %v \nWriter test %v\nPrinter test1 %v\nWritePointer address %v", &jj, &test, &test1, &test2)

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hi This is my interface sample by 8 per line and lets see how it goes"))
	wc.Close()
}

type WritePointer interface {
	WriteObjPointer(string)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface { // compose interface
	Writer
	Closer
}

type Printer interface {
	Print()
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (feed ConsoleWriter) Print() {
	fmt.Printf("\nPrint called %v", feed.name)
}

type ConsoleWriter struct {
	name string
	age  int16
}

func (feed ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
func (feed *ConsoleWriter) WriteObjPointer(text string) {
	feed.name = text
	fmt.Printf("\n%v", feed)
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
