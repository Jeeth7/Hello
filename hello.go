package main

import (
	"fmt"
)

var (
	i int16 = 10
	j int   = 3
)

const (
	isAdmin = 1 << iota
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
}
