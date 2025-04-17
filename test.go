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
	A.Sort()
	fmt.Println("has in A empty", A)

	B, err := set.NewInt8Set([]int8{5, 4, 2, 9}...)
	if err != nil {
		fmt.Println("B err", err)
	} else {
		fmt.Println("B", B)
	}

	if err := A.Add(8); err != nil {
		fmt.Println("A Add err", err)
	} else {
		fmt.Println("A Add", A)
	}

	if err := B.Add(11); err != nil {
		fmt.Println("B Add err", err)
	} else {
		fmt.Println("B Add", B)
	}

	A.Add(20)
	A.Add(100)

	fmt.Println("set B", B)
	ns := A.Discard(9)
	fmt.Println("set A", ns)
	diff, _ := B.Difference(ns)
	fmt.Println("B difference A ", diff)
	inter, _ := B.Intersect(ns)
	fmt.Println("B intersect A ", inter)

	B.Sort()
	fmt.Println("B sort", B)
	B.ReverseSort()
	fmt.Println("B reverse", B)

	fmt.Println("B has", B.Has(8))
	fmt.Println("B has", B.Has(10))
	slice, _ := B.ToSlice()
	fmt.Println("B slice", slice)
	unione, err := B.Union(A)
	fmt.Println("B unione", unione, err)

}
