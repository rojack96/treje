package main

import (
	"fmt"

	set "github.com/rojack96/treje/set"
)

// -------------------------------------------------------------

func main() {
	A, err := set.NewInt8Set()
	if err != nil {
		fmt.Println("A err", err)
	} else {
		fmt.Println("A", A)
	}

	B, err := set.NewInt8Set([]int8{5, 4, 2, 9}...)
	if err != nil {
		fmt.Println("B err", err)
	} else {
		fmt.Println("B", B)
	}

	B.Discard(9)
	fmt.Println("B unione", B, len(B))

}
