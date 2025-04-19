package main

import (
	"fmt"
	mapSet "github.com/rojack96/treje/map_set"
)

// -------------------------------------------------------------

func main() {
	A, err := mapSet.NewInt8Set()
	if err != nil {
		fmt.Println("A err", err)
	} else {
		fmt.Println("A", A)
	}

	B, err := mapSet.NewInt8Set([]int8{5, 4, 2, 9}...)
	if err != nil {
		fmt.Println("B err", err)
	} else {
		fmt.Println("B", B)
	}

	test := B.Equals(A)

	fmt.Println("B unione", test, len(B))

}
