package gopractice

import (
	"fmt"
	"math"
)

var (
	i int16 = 10
	j int   = 3
)

const (
	isAdmin = 1 << iota //starts with 0 and increase with 1 in subsequent values
	isHeadquarters
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	k := "THis is a string"
	j := 'f'
	var l byte = 'l'
	const tes = 53

	b := []byte(k)
	b[3] = byte('d')
	b[1] = l
	fmt.Printf("%v, %T\n", string(b), b[3])
	fmt.Printf("%v, %T\n", l, l)
	fmt.Printf("%v, %T\n", j, j)

	fmt.Printf("%v, %T\n", tes, tes)
	fmt.Printf("%v, %T\n", tes+i, tes+i)

	//Check roles from const iota << 1
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("isAdmin %b, canSeeFinancials %b, canSeeEurope %b\n", isAdmin, canSeeFinancials, canSeeEurope)
	fmt.Printf("roles %b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)

	//ARRay declaration

	grades1 := [...]int{84, 22, 44, 93, 90}
	fmt.Printf("grades1 %v %T\n", grades1, grades1)
	grades2 := [...]int{0, 1, 2, 3, 5, 8}
	fmt.Printf("grades2 %v %T\n", grades2, grades2)
	fmt.Printf("value %v length %v capacity %v\n", grades2, len(grades2), cap(grades2))
	grades3 := [5]int{84, 22, 44, 93, 90}
	fmt.Printf("grades3 %v %T\n", grades3, grades3)
	grades4 := [...]int{84, 22, 44, 93, 90}
	fmt.Printf("grades4 %v %v\n", grades4, grades4[2:4])
	grades6 := append(grades2[:], 22)
	fmt.Printf("value %v length %v capacity %v type %T\n", grades6, len(grades6), cap(grades6), grades6)

	grades7 := make([]int, 3, 59)
	grades7[0] = 1
	grades7[1] = 1
	grades7[2] = 1
	fmt.Printf("value %v length %v capacity %v\n", grades7, len(grades7), cap(grades7))
	grades8 := append(grades7, 22, 26, 78, 45)
	grades8[2] = 56
	fmt.Printf("value %v length %v capacity %v\n", grades7, len(grades7), cap(grades7))
	fmt.Printf("value %v length %v capacity %v\n", grades8, len(grades8), cap(grades8))
	grades9 := append(grades8[2:3], grades8[4:5]...)
	fmt.Printf("value %v length %v capacity %v\n", grades9, len(grades9), cap(grades9))
	fmt.Printf("value %v length %v capacity %v\n", grades8, len(grades8), cap(grades8))
	//MAP

	statePopulation := map[string]int{
		"TamilNadu": 123456,
		"Kerala":    236583,
		"Jharkhand": 859674,
	}
	fmt.Println(statePopulation)
	statePopulation["Odisha"] = 258963   //Add
	delete(statePopulation, "Jharkhand") //delete
	fmt.Println(statePopulation)

	cityPopulation := make(map[string]int)
	cityPopulation = map[string]int{
		"Delhi":   5585,
		"Chennai": 5589,
	}
	fmt.Println(cityPopulation)
	fmt.Println(cityPopulation["Trivandrum"])
	tvm, ok := cityPopulation["Trivandrum"]
	_, ok1 := cityPopulation["Delhi"]
	fmt.Println(tvm, ok)
	fmt.Println(ok1)

	chiefDoctor := Doctor{
		name:   "Peter",
		number: 589321,
		companion: []string{
			"jack",
			"jill",
		},
	}

	fmt.Println(chiefDoctor)

	localDoctor := struct{ name string }{name: "Jack Ma"}
	anotherlocalDoctory := localDoctor
	referencelocalDoctory := &localDoctor
	fmt.Print("*************")
	defer fmt.Printf("\n*******%v", localDoctor) //defer will print execute the statement after the main is returned
	localDoctor.name = "Jack Jon"
	fmt.Println(localDoctor)
	fmt.Println(anotherlocalDoctory)
	defer fmt.Println(referencelocalDoctory)
	defer fmt.Println(print())

	if true {
		fmt.Println("This is always true")
	}

	if val, ok := statePopulation["Odisha"]; ok {
		fmt.Printf("Population is %v", val)
	} else {
		println("Not found")
	}

	println(j)
	val := 0.1323
	if val == math.Pow(math.Sqrt(val), 2) {
		println("the values are same")
	} else {
		println("the values are different")
	}

	//comparison on floating
	val1 := 0.1323
	if (math.Abs(val1)/math.Pow(math.Sqrt(val1), 2) - 1) < 0.001 {
		println("the values are same")
	} else {
		println("the values are different")
	}

	//Swith with tags
	switch 10 {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 0:
		println("Single digit number")
	case 10:
		println("Ten")
	default:
		println("Some other number")

	}

	//Switch without tags
	i := 106
	switch {
	case i < 150:
		println("less than 150")
		fallthrough //optional keyword to execute next without evaluating condition
	case i < 100:
		println("Less than 100")
	}

	//switch statement with type
	var test interface{} = 0.99
	switch test.(type) {
	case int:
		fmt.Print("This is integer")
		if test == 0.99 {
			break
		}
		fmt.Println("this wil be breaked")
	case float32, float64:
		fmt.Print("this is float")
		if test == 1.99 {
			break
		}
		fmt.Println("this wil be breaked")
	case string:
		fmt.Println("this is a string")
	default:
		fmt.Println("This is some other type")
	}

	for i, j := 0, 9.9; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	i = 0
	for i < 5 { // same as for ;i < 5; {}
		fmt.Println(i)
		i++
	}
	fmt.Println("----------------------------------------------------------------------------")
LOOP:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			//if i%2 == 0 {
			//	continue
			//}
			if i == 3 {
				break LOOP
			}
			fmt.Println(i * j)
		}
	}

	fmt.Println("----------------------------------------------------------------------------")

	for k, v := range grades1 {
		fmt.Println(k, v)
	}

	for k, v := range cityPopulation {
		fmt.Println(k, v)
	}
}

type Animal struct {
	name string
}

type Doctor struct {
	name      string
	number    int
	companion []string
}

func print() string {
	fmt.Printf("****I am called")
	return "*called"
}
