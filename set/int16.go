package set

import "strconv"

type Int16 int16

type Int16Set []Int16

func (set *Int16Set) Add(elem Int16) {
	for _, n := range *set {
		if (n ^ elem) == 0 {
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
