package set

import (
	"strconv"
)

type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Int int

type Int8Set []Int8
type Int16Set []Int16
type Int32Set []Int32
type Int64Set []Int64
type IntSet []Int

/* ---------- Int8 ---------- */

func (set *Int8Set) Add(elem Int8) Int8Set {
	for _, n := range *set {
		equal := (n ^ elem) == 0
		if !equal {
			continue
		}
		panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return *set
}

func (set *Int8Set) ToSlice() []int8 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]int8, len(*set))
	for i, v := range *set {
		result[i] = int8(v)
	}
	return result
}

/* ---------- Int16 ---------- */

func (set *Int16Set) Add(elem Int16) Int16Set {
	for _, n := range *set {
		equal := (n ^ elem) == 0
		if !equal {
			continue
		}
		panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return *set
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

/* ---------- Int32 ---------- */

func (set *Int32Set) Add(elem Int32) Int32Set {
	for _, n := range *set {
		equal := (n ^ elem) == 0
		if !equal {
			continue
		}
		panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return *set
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

/* ---------- Int64 ---------- */

func (set *Int64Set) Add(elem Int64) Int64Set {
	for _, n := range *set {
		equal := (n ^ elem) == 0
		if !equal {
			continue
		}
		panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return *set
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

/* ---------- Int ---------- */

func (set *IntSet) Add(elem Int) IntSet {
	for _, n := range *set {
		equal := (n ^ elem) == 0
		if !equal {
			continue
		}
		panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return *set
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
