package main

import (
	"fmt"

	set "github.com/rojack96/treje/set"
)

// -------------------------------------------------------------

func main() {
	t := []int8{5, 4, 2, 9}
	d := []int8{5, 4, 7, 10}
	s := set.NewInt8Set(t...)
	p := set.NewInt8Set(d...)
	s.Sort()
	//x := s.Difference(p)
	y := p.Difference(s)
	//s.Sort()
	fmt.Println("s", y)
}
