package set

import "strconv"

type Int int

type IntSet []Int

func (set *IntSet) Add(elem Int) {
	for _, n := range *set {
		if (n ^ elem) == 0 {
			panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *IntSet) ToSlice() []int {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]int, len(*set))
	for i, v := range *set {
		result[i] = int(v)
	}
	return result
}
