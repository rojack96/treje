package set

import (
	"errors"
	"sort"
	"strconv"
)

type Int8 *int8
type Int8Set []Int8

// NewInt8Set - Create a new empty set or from a slice
func NewInt8Set(elems ...int8) (Int8Set, error) {
	set := Int8Set{}

	if len(elems) == 0 {
		return set, nil
	}

	if len(elems) == 1 {
		return append(set, &elems[0]), nil
	}

	elemsCopy := make([]int8, len(elems))
	copy(elemsCopy, elems)

	sort.Slice(elems, func(i, j int) bool {
		return elems[i] < elems[j]
	})

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			return nil, errors.New(HasDuplicates)
		}
	}

	for _, n := range elemsCopy {
		set = append(set, &n)
	}

	return set, nil
}

// Has - Return true if element is in set, otherwise false
func (set *Int8Set) Has(elem Int8) bool {
	for _, n := range *set {
		if n == elem {
			return true
		}
	}
	return false
}

// Add - Append a new element to the set if and only if it is not already present
func (set *Int8Set) Add(elem Int8) error {
	if set.Has(elem) {
		return errors.New(strconv.Itoa(int(*elem)) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return nil
}

// Remove - Remove a specific element from set, if element not exist raise an error
func (set *Int8Set) Remove(elem Int8) error {
	originalLen := len(*set)
	set.Discard(elem)
	if len(*set) == originalLen {
		return errors.New(ElemNotExist)
	}
	return nil
}

// Discard - Remove a specific element from set
func (set *Int8Set) Discard(elem Int8) Int8Set {
	result := *set
	for i, n := range result {
		if n == elem {
			*set = append(result[:i], result[i+1:]...)
		}
	}

	return *set
}

// Pop - Remove and return element from set at a given index (or last if none provided)
func (set *Int8Set) Pop(index ...int) (int8, error) {
	if len(*set) == 0 {
		return 0, errors.New(EmptySet)
	}

	i := len(*set) - 1
	if len(index) > 0 {
		i = index[0]
		if i < 0 || i >= len(*set) {
			return 0, errors.New("index out of range")
		}
	}

	elem := (*set)[i]
	*set = append((*set)[:i], (*set)[i+1:]...)
	return *elem, nil
}

func (set *Int8Set) Union(b Int8Set) Int8Set {
	*set = append(*set, b...)
	return *set
}

// Intersect - Returns the elements that are present in both input sets.
func (set *Int8Set) Intersect(b Int8Set) Int8Set {
	set.Sort()
	b.Sort()

	var result Int8Set
	i, j := 0, 0

	for i < len(*set) && j < len(b) {
		if (*set)[i] == b[j] {
			if len(result) == 0 || result[len(result)-1] != (*set)[i] {
				result = append(result, (*set)[i])
			}
			i++
			j++
		} else if *(*set)[i] < *b[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

// Difference - Returns the elements that are present in the first set
// but not in the second set.
func (set *Int8Set) Difference(b Int8Set) Int8Set {
	var result Int8Set

	for _, elemA := range *set {
		found := false
		for _, elemB := range b {
			if elemA == elemB {
				found = true
				break
			}
		}
		if !found {
			result = append(result, elemA)
		}
	}

	return result
}

// Clear - Remove all element
func (set *Int8Set) Clear() {
	*set = Int8Set{}
}

// Min - Return minimum element from set
func (set *Int8Set) Min() int8 {
	minimum := *set
	minimum.Sort()

	res := minimum[0]
	return *res
}

// Max - Return maximum element from set
func (set *Int8Set) Max() int8 {
	maximum := *set
	maximum.Sort()

	res := maximum[len(maximum)-1]
	return *res
}

// Sum - Return a sum of all elements
func (set *Int8Set) Sum() int {
	total := 0

	if len(*set) > 0 {
		for _, v := range *set {
			total += int(*v)
		}
	}

	return total
}

// Sort - Sort element in ascending mode
func (set *Int8Set) Sort() {
	sort.Slice(*set, func(i, j int) bool {
		return *(*set)[i] < *(*set)[j]
	})
}

// ReverseSort - Sort element in descending mode
func (set *Int8Set) ReverseSort() {
	sort.Slice(*set, func(i, j int) bool {
		return *(*set)[i] > *(*set)[j]
	})
}

// ToSlice - Returns a slice of native datatype from the set
func (set *Int8Set) ToSlice() []int8 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]int8, len(*set))
	for i, v := range *set {
		result[i] = *v
	}
	return result
}
