package main

import "fmt"

func main() {

	func() {
		fmt.Println("this is an anonymous functtion")
	}()

	testFunc := func() {
		fmt.Println("this is an anonymous functtion")
	}
	testFunc()
	var testFunc2 func(int, int, string) (string, error)
	testFunc2 = func(n, d int, name string) (string, error) {
		if d == 0 {
			return name, fmt.Errorf("denominator cannot be zero!")
		} else {
			return name, nil
		}
	}
	st, er := testFunc2(2, 0, "test")
	fmt.Println("\n", st, er)
	test("parameter 1")
	wor1 := "his car"
	wor2 := "his mobile"
	fmt.Printf("\n%v, %v", wor1, wor2)
	marriage(&wor1, &wor2)
	fmt.Printf("\n%v, %v", wor1, wor2)
	result := variaticParameter(1, 2, 1, 3, 3, 2)
	fmt.Printf("\nResult sum = %v", *result)
	result2 := namedReturnType(1, 2, 3, 4, 5)
	fmt.Printf("\nResult sum = %v", result2)

	r3, e1 := multipleReturnVariable(4, 0)
	fmt.Printf("\nResult sum = %v with error = %v", r3, e1)

	//Playing with methods
	inderjit := people{
		name: "Inderjit",
	}
	fmt.Println(inderjit)
	inderjit.addAge(33)
	fmt.Println(inderjit)

	inderjit.addAgePointer(33)
	fmt.Println(inderjit)
}

func test(t string) {
	fmt.Printf("This is the function with a parameter %v", t)
}

func marriage(t, j *string) {
	*t = "her car"
	*j = "her mobile"
}

func variaticParameter(value ...int) *int {
	fmt.Println(value)
	var sum int
	for _, v := range value {
		sum = sum + v
	}
	return &sum // it is safe to return pointers and this memory willnot be destroyed after this function
}

func namedReturnType(value ...int) (sum int) {
	fmt.Println(value)
	//var sum int this line is not needed since it is already declared
	for _, v := range value {
		sum = sum + v
	}
	return //sum only return keyword is needed
}

func multipleReturnVariable(a, b float32) (float32, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

type people struct {
	name string
	age  int
}

func (person people) addAge(val int) {
	person.age = val
}
func (person *people) addAgePointer(val int) {
	person.age = val
}
