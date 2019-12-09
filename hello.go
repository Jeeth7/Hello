package main

import (
	"fmt"
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
	localDoctor.name = "Jack Jon"
	fmt.Println(localDoctor)
	fmt.Println(anotherlocalDoctory)
	fmt.Println(referencelocalDoctory)

}

type Animal struct{
	name string
	
}

type Doctor struct {
	name      string
	number    int
	companion []string
}

//func print(val){
//	fmt.Printf("%v %T length %v capacity %v ", val, val, len(val), cap(val))
//}
