package set

import (
	"sort"
	"strconv"
)

type Uint32 uint32
type Uint32Set []Uint32

func NewUint32Set(elems ...uint32) Uint32Set {
	set := Uint32Set{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Uint32(elems[0]))
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
		set = append(set, Uint32(n))
	}

	return set
}

func (set *Uint32Set) Add(elem Uint32) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Uint32Set) ToSlice() []uint32 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]uint32, len(*set))
	for i, v := range *set {
		result[i] = uint32(v)
	}
	return result
}
