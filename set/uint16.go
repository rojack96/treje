package set

import "strconv"

type Uint16 uint16
type Uint16Set []Uint16

func (set *Uint16Set) Add(elem Uint16) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Uint16Set) ToSlice() []uint16 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]uint16, len(*set))
	for i, v := range *set {
		result[i] = uint16(v)
	}
	return result
}
