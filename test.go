package main

import (
	"fmt"
	"github.com/rojack96/treje/set"
)

// -------------------------------------------------------------

func main() {
	A, err := set.New().Int8()
	if err != nil {
		fmt.Println("A err", err)
	} else {
		fmt.Println("A", A)
	}

	B, err := set.New().Int8([]int8{5, 4, 2, 9}...)
	if err != nil {
		fmt.Println("B err", err)
	} else {
		fmt.Println("B", B)
	}

	test := B.Equals(A)

	fmt.Println("B unione", test, len(B))

}
