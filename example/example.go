package main

import (
	"fmt"
	"github.com/rojack96/treje"
)

// -------------------------------------------------------------

func main() {

	A, err := treje.NewMapSet().Int8([]int8{5}...)
	if err != nil {
		fmt.Println("A err", err)
	}

	B, err := treje.NewMapSet().Int8([]int8{5, 4, 2, 9}...)
	if err != nil {
		fmt.Println("B err", err)
	}

	test := A.Intersect(B)

	fmt.Println("test", test)

}
