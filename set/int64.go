package set

import (
	"sort"
	"strconv"
)

type Int64 int64
type Int64Set []Int64

func NewInt64Set(elems ...int64) Int64Set {
	set := Int64Set{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Int64(elems[0]))
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
		set = append(set, Int64(n))
	}

	return set
}

func (set *Int64Set) Add(elem Int64) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Int64Set) ToSlice() []int64 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]int64, len(*set))
	for i, v := range *set {
		result[i] = int64(v)
	}
	return result
}
