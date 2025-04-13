package set

import "strconv"

type Uint8 uint8
type Uint8Set []Uint8

func (set *Uint8Set) Add(elem Uint8) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Uint8Set) ToSlice() []uint8 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]uint8, len(*set))
	for i, v := range *set {
		result[i] = uint8(v)
	}
	return result
}
