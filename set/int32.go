package set

import (
	"sort"
	"strconv"
)

type Int32 int32
type Int32Set []Int32

func NewInt32Set(elems ...int32) Int32Set {
	set := Int32Set{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Int32(elems[0]))
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
		set = append(set, Int32(n))
	}

	return set
}

func (set *Int32Set) Add(elem Int32) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Int32Set) ToSlice() []int32 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]int32, len(*set))
	for i, v := range *set {
		result[i] = int32(v)
	}
	return result
}
