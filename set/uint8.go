package set

import (
	"sort"
	"strconv"
)

type Uint8 uint8
type Uint8Set []Uint8

func NewUint8Set(elems ...uint8) Uint8Set {
	set := Uint8Set{}

	if len(elems) == 0 {
		return Uint8Set{}
	}

	if len(elems) == 1 {
		set = append(set, Uint8(elems[0]))
		return set
	}

	sort.Slice(elems, func(i, j int) bool {
		return elems[i] < elems[j]
	})

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elems {
		set = append(set, Uint8(n))
	}

	return set
}

func (set *Uint8Set) Add(elem Uint8) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Uint8Set) ToSlice() []uint8 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]uint8, len(*set))
	for i, v := range *set {
		result[i] = uint8(v)
	}
	return result
}
