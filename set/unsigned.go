package set

import (
	"strconv"
)

type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64
type Uint uint

type Uint8Set []Uint8
type Uint16Set []Uint16
type Uint32Set []Uint32
type Uint64Set []Uint64
type UintSet []Uint

/* ---------- Uint8 ---------- */

func (set *Uint8Set) Add(elem Uint8) Uint8Set {
	for _, n := range *set {
		equal := (n ^ elem) == 0
		if !equal {
			continue
		}
		panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return *set
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

/* ---------- Uint16 ---------- */

func (set *Uint16Set) Add(elem Uint16) Uint16Set {
	for _, n := range *set {
		equal := (n ^ elem) == 0
		if !equal {
			continue
		}
		panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return *set
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

/* ---------- Uint32 ---------- */

func (set *Uint32Set) Add(elem Uint32) Uint32Set {
	for _, n := range *set {
		equal := (n ^ elem) == 0
		if !equal {
			continue
		}
		panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return *set
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

/* ---------- Uint64 ---------- */

func (set *Uint64Set) Add(elem Uint64) Uint64Set {
	for _, n := range *set {
		equal := (n ^ elem) == 0
		if !equal {
			continue
		}
		panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return *set
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

/* ---------- Uint ---------- */

func (set *UintSet) Add(elem Uint) UintSet {
	for _, n := range *set {
		equal := (n ^ elem) == 0
		if !equal {
			continue
		}
		panic(strconv.FormatUint(uint64(elem), 10) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return *set
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
