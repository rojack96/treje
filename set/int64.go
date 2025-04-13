package set

import "strconv"

type Int64 int64
type Int64Set []Int64

func (set *Int64Set) Add(elem Int64) {
	for _, n := range *set {
		if (n ^ elem) == 0 {
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
