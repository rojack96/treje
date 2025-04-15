package set

import (
	"sort"
	"strconv"
)

type Int16 int16

type Int16Set []Int16

func NewInt16Set(elems ...int16) Int16Set {
	set := Int16Set{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Int16(elems[0]))
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
		set = append(set, Int16(n))
	}

	return set
}

func (set *Int16Set) Add(elem Int16) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Int16Set) ToSlice() []int16 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]int16, len(*set))
	for i, v := range *set {
		result[i] = int16(v)
	}
	return result
}
