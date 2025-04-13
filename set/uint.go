package set

import "strconv"

type Uint uint
type UintSet []Uint

func (set *UintSet) Add(elem Uint) {
	for _, n := range *set {
		if (n ^ elem) == 0 {
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
