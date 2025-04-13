package set

import "strconv"

type Int32 int32
type Int32Set []Int32

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
