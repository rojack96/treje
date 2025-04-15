package set

import (
	"sort"
	"strconv"
)

type Uint64 uint64
type Uint64Set []Uint64

func NewUint64Set(elems ...uint64) Uint64Set {
	set := Uint64Set{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Uint64(elems[0]))
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
		set = append(set, Uint64(n))
	}

	return set
}

func (set *Uint64Set) Add(elem Uint64) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Uint64Set) ToSlice() []uint64 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]uint64, len(*set))
	for i, v := range *set {
		result[i] = uint64(v)
	}
	return result
}
