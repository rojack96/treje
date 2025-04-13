package main

import (
	"fmt"

	set "github.com/rojack96/treje/set"
)

// -------------------------------------------------------------

func main() {
	s := set.NewInt8Set([]int8{2, 4, 5, 4, 2}...)

	fmt.Println("s", s)
}
