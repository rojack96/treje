package set

import "strconv"

type Uint32 uint32
type Uint32Set []Uint32

func (set *Uint32Set) Add(elem Uint32) {
	for _, n := range *set {
		if (n ^ elem) == 0 {
			panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Uint32Set) ToSlice() []uint32 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]uint32, len(*set))
	for i, v := range *set {
		result[i] = uint32(v)
	}
	return result
}
