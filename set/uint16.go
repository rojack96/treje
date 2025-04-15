package set

import (
	"sort"
	"strconv"
)

type Uint16 uint16
type Uint16Set []Uint16

func NewUint16Set(elems ...uint16) Uint16Set {
	set := Uint16Set{}

	if len(elems) == 0 {
		return Uint16Set{}
	}

	if len(elems) == 1 {
		set = append(set, Uint16(elems[0]))
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
		set = append(set, Uint16(n))
	}

	return set
}

func (set *Uint16Set) Add(elem Uint16) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Uint16Set) ToSlice() []uint16 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]uint16, len(*set))
	for i, v := range *set {
		result[i] = uint16(v)
	}
	return result
}
