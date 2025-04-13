package main

import (
	"fmt"

	set "github.com/rojack96/treje/set"
)

// -------------------------------------------------------------

func main() {
	t := []int8{5, 4, 2, 9}
	s := set.NewInt8Set(t...)
	//s.Sort()
	fmt.Println("s", s)
}
