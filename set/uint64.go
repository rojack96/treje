package set

import "strconv"

type Uint64 uint64
type Uint64Set []Uint64

func (set *Uint64Set) Add(elem Uint64) {
	for _, n := range *set {
		if (n ^ elem) == 0 {
			panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Uint64Set) ToSlice() []uint64 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]uint64, len(*set))
	for i, v := range *set {
		result[i] = uint64(v)
	}
	return result
}
