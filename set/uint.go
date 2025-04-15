package set

import (
	"sort"
	"strconv"
)

type Uint uint
type UintSet []Uint

func NewUintSet(elems ...uint) UintSet {
	set := UintSet{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Uint(elems[0]))
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
		set = append(set, Uint(n))
	}

	return set
}

func (set *UintSet) Add(elem Uint) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *UintSet) ToSlice() []uint {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]uint, len(*set))
	for i, v := range *set {
		result[i] = uint(v)
	}
	return result
}
