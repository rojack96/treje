package set

import (
	"sort"
)

/* ---------- Integers ---------- */

func NewInt8Set(elems ...int8) Int8Set {
	set := Int8Set{}
	elemsCopy := make([]int8, len(elems))
	copy(elemsCopy, elems)

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Int8(elems[0]))
		return set
	}

	quickSort8(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elemsCopy {
		set = append(set, Int8(n))
	}

	return set
}

func NewInt16Set(elems ...int16) Int16Set {
	set := Int16Set{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Int16(elems[0]))
		return set
	}

	quickSort16(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elems {
		set = append(set, Int16(n))
	}

	return set
}

func NewInt32Set(elems ...int32) Int32Set {
	set := Int32Set{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Int32(elems[0]))
		return set
	}

	quickSort32(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elems {
		set = append(set, Int32(n))
	}

	return set
}

func NewInt64Set(elems ...int64) Int64Set {
	set := Int64Set{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Int64(elems[0]))
		return set
	}

	quickSort64(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elems {
		set = append(set, Int64(n))
	}

	return set
}

func NewIntSet(elems ...int) IntSet {
	set := IntSet{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Int(elems[0]))
		return set
	}

	sort.Ints(elems)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
		set = append(set, Int(elems[i]))
	}

	for _, n := range elems {
		set = append(set, Int(n))
	}

	return set
}

/* ---------- Unsigned Integers ---------- */

func NewUint8Set(elems ...uint8) Uint8Set {
	set := Uint8Set{}

	if len(elems) == 0 {
		return Uint8Set{}
	}

	if len(elems) == 1 {
		set = append(set, Uint8(elems[0]))
		return set
	}

	quickSortUnsigned8(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elems {
		set = append(set, Uint8(n))
	}

	return set
}

func NewUint16Set(elems ...uint16) Uint16Set {
	set := Uint16Set{}

	if len(elems) == 0 {
		return Uint16Set{}
	}

	if len(elems) == 1 {
		set = append(set, Uint16(elems[0]))
		return set
	}

	quickSortUnsigned16(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elems {
		set = append(set, Uint16(n))
	}

	return set
}

func NewUint32Set(elems ...uint32) Uint32Set {
	set := Uint32Set{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Uint32(elems[0]))
		return set
	}

	quickSortUnsigned32(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elems {
		set = append(set, Uint32(n))
	}

	return set
}

func NewUint64Set(elems ...uint64) Uint64Set {
	set := Uint64Set{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Uint64(elems[0]))
		return set
	}

	quickSortUnsigned64(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elems {
		set = append(set, Uint64(n))
	}

	return set
}

func NewUintSet(elems ...uint) UintSet {
	set := UintSet{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Uint(elems[0]))
		return set
	}

	quickSortUnsigned(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elems {
		set = append(set, Uint(n))
	}

	return set
}

func NewStringSet(elems ...string) StringSet {
	set := StringSet{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, String(elems[0]))
		return set
	}

	quickSortString(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elems {
		set = append(set, String(n))
	}

	return set
}

/* ---------- Sort functions ---------- */

/* ---------- Quick Sort Int8 ---------- */
func partition8(slice []int8, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func quickSort8(slice []int8, low, high int) {
	if low < high {
		pivot := partition8(slice, low, high)
		quickSort8(slice, low, pivot-1)
		quickSort8(slice, pivot+1, high)
	}
}

/* ---------- Quick Sort Int16 ---------- */

func partition16(slice []int16, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func quickSort16(slice []int16, low, high int) {
	if low < high {
		pivot := partition16(slice, low, high)
		quickSort16(slice, low, pivot-1)
		quickSort16(slice, pivot+1, high)
	}
}

/* ---------- Quick Sort Int32 ---------- */

func partition32(slice []int32, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func quickSort32(slice []int32, low, high int) {
	if low < high {
		pivot := partition32(slice, low, high)
		quickSort32(slice, low, pivot-1)
		quickSort32(slice, pivot+1, high)
	}
}

/* ---------- Quick Sort Int64 ---------- */

func partition64(slice []int64, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func quickSort64(slice []int64, low, high int) {
	if low < high {
		pivot := partition64(slice, low, high)
		quickSort64(slice, low, pivot-1)
		quickSort64(slice, pivot+1, high)
	}
}

/* ---------- Quick Sort Uint8 ---------- */

func partitionUnsigned8(slice []uint8, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func quickSortUnsigned8(slice []uint8, low, high int) {
	if low < high {

		pivot := partitionUnsigned8(slice, low, high)

		quickSortUnsigned8(slice, low, pivot-1)
		quickSortUnsigned8(slice, pivot+1, high)
	}
}

/* ---------- Quick Sort Uint16 ---------- */

func partitionUnsigned16(slice []uint16, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func quickSortUnsigned16(slice []uint16, low, high int) {
	if low < high {

		pivot := partitionUnsigned16(slice, low, high)

		quickSortUnsigned16(slice, low, pivot-1)
		quickSortUnsigned16(slice, pivot+1, high)
	}
}

/* ---------- Quick Sort Uint32 ---------- */

func partitionUnsigned32(slice []uint32, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func quickSortUnsigned32(slice []uint32, low, high int) {
	if low < high {
		pivot := partitionUnsigned32(slice, low, high)
		quickSortUnsigned32(slice, low, pivot-1)
		quickSortUnsigned32(slice, pivot+1, high)
	}
}

/* ---------- Quick Sort Uint64 ---------- */

func partitionUnsigned64(slice []uint64, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func quickSortUnsigned64(slice []uint64, low, high int) {
	if low < high {
		pivot := partitionUnsigned64(slice, low, high)
		quickSortUnsigned64(slice, low, pivot-1)
		quickSortUnsigned64(slice, pivot+1, high)
	}
}

/* ---------- Quick Sort Uint ---------- */

func partitionUnsigned(slice []uint, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func quickSortUnsigned(slice []uint, low, high int) {
	if low < high {
		pivot := partitionUnsigned(slice, low, high)
		quickSortUnsigned(slice, low, pivot-1)
		quickSortUnsigned(slice, pivot+1, high)
	}
}

/* ---------- Quick Sort String ---------- */

func partitionString(slice []string, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func quickSortString(slice []string, low, high int) {
	if low < high {
		pivot := partitionString(slice, low, high)
		quickSortString(slice, low, pivot-1)
		quickSortString(slice, pivot+1, high)
	}
}
