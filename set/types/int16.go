package types

import (
	"errors"
	"github.com/rojack96/treje/common"
	"sort"
	"strconv"
)

type (
	Integer16 int16
	Int16Set  []Integer16
)

// Int16 - Create a new empty set or from a slice
func (s Set) Int16(elems ...int16) (Int16Set, error) {
	set := Int16Set{}

	if len(elems) == 0 {
		return set, nil
	}

	if len(elems) == 1 {
		return append(set, Integer16(elems[0])), nil
	}

	elemsCopy := make([]int16, len(elems))
	copy(elemsCopy, elems)

	sort.Slice(elems, func(i, j int) bool {
		return elems[i] < elems[j]
	})

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			return nil, errors.New(common.HasDuplicates)
		}
	}

	for _, n := range elemsCopy {
		set = append(set, Integer16(n))
	}

	return set, nil
}

/*
	Manipulation set methods
*/

// Add - Append a new element to the set if and only if it is not already present
func (set *Int16Set) Add(elem Integer16) error {
	if set.Has(elem) {
		return errors.New(strconv.Itoa(int(elem)) + " " + common.AlreadyExists)
	}

	*set = append(*set, elem)
	return nil
}

// Remove - Remove a specific element from a set, if the element not exists raise an error
func (set *Int16Set) Remove(elem Integer16) error {
	if set.IsEmpty() {
		return errors.New(common.EmptySet)
	}

	originalLen := len(*set)
	set.Discard(elem)
	if len(*set) == originalLen {
		return errors.New(common.ElemNotExist)
	}
	return nil
}

// Discard - Remove a specific element from set
func (set *Int16Set) Discard(elem Integer16) {
	result := *set
	for i, n := range result {
		if n == elem {
			*set = append(result[:i], result[i+1:]...)
			break
		}
	}
}

// Pop - Remove and return element from a set at a given index (or last if none provided)
func (set *Int16Set) Pop(index ...int) (int16, error) {
	if set.IsEmpty() {
		return 0, errors.New(common.EmptySet)
	}

	i := len(*set) - 1
	if len(index) > 0 {
		i = index[0]
		if i < 0 || i >= len(*set) {
			return 0, errors.New(common.IndexOutOfRange)
		}
	}

	elem := (*set)[i]
	*set = append((*set)[:i], (*set)[i+1:]...)
	return int16(elem), nil
}

/*
	Set operation methods
*/

// Union - Merges the current set with another set, but returns an error
// if there are any duplicates in the union.
func (set *Int16Set) Union(b Int16Set) (Int16Set, error) {

	for _, elemB := range b {
		if set.Has(elemB) {
			return *set, errors.New(common.HasDuplicates)
		}
		*set = append(*set, elemB)
	}
	return *set, nil
}

// Intersect - Returns the elements that are present in both input sets.
func (set *Int16Set) Intersect(b Int16Set) (Int16Set, error) {
	set.Sort()
	b.Sort()

	var result Int16Set
	i, j := 0, 0

	for i < len(*set) && j < len(b) {
		if (*set)[i] == b[j] {
			if len(result) == 0 || result[len(result)-1] != (*set)[i] {
				result = append(result, (*set)[i])
			}
			i++
			j++
		} else if (*set)[i] < b[j] {
			i++
		} else {
			j++
		}
	}

	return result, nil
}

// Difference - Returns the elements that are present in the first set
// but not in the second set.
func (set *Int16Set) Difference(b Int16Set) (Int16Set, error) {
	var result Int16Set

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

	return result, nil
}

// SymmetricDifference - Returns a new set with elements that are present in either of the two sets but not in both.
func (set *Int16Set) SymmetricDifference(b Int16Set) (Int16Set, error) {
	var (
		diff1, diff2 Int16Set
		err          error
	)
	if diff1, err = set.Difference(b); err != nil {
		return nil, err
	}
	if diff2, err = (&b).Difference(*set); err != nil {
		return nil, err
	}

	return append(diff1, diff2...), nil
}

// IsSubsetOf - Returns true if the current set is a subset of the given set b.
func (set *Int16Set) IsSubsetOf(b Int16Set) bool {
	for _, elem := range *set {
		found := false
		for _, other := range b {
			if elem == other {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Equals - Returns true if the current set and set b contain the same elements.
func (set *Int16Set) Equals(b Int16Set) bool {
	return set.IsSubsetOf(b) && (&b).IsSubsetOf(*set)
}

/*
	Utility methods
*/

// Has - Return true if the element is in set, otherwise false
func (set *Int16Set) Has(elem Integer16) bool {
	for _, n := range *set {
		if n == elem {
			return true
		}
	}
	return false
}

// IsEmpty - Return true if the set is empty, else false
func (set *Int16Set) IsEmpty() bool {
	return len(*set) == 0
}

// Clear - Remove all elements
func (set *Int16Set) Clear() {
	*set = Int16Set{}
}

// Min - Return minimum element from the set
func (set *Int16Set) Min() int16 {
	if set.IsEmpty() {
		return 0
	}

	minimum := *set
	minimum.Sort()

	res := minimum[0]
	return int16(res)
}

// Max - Return maximum element from the set
func (set *Int16Set) Max() int16 {
	if set.IsEmpty() {
		return 0
	}

	maximum := *set
	maximum.Sort()

	res := maximum[len(maximum)-1]
	return int16(res)
}

// Sum - Return a sum of all elements
func (set *Int16Set) Sum() int {
	total := 0

	if len(*set) > 0 {
		for _, v := range *set {
			total += int(v)
		}
	}

	return total
}

// Sort - Sort element in ascending mode
func (set *Int16Set) Sort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] < (*set)[j]
	})
}

// ReverseSort - Sort element in descending mode
func (set *Int16Set) ReverseSort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] > (*set)[j]
	})
}

/*
	Methods to manipulate a set object
*/

func (set *Int16Set) Copy() (Int16Set, error) {
	if set.IsEmpty() {
		return nil, errors.New("cannot copy an empty slice")
	}
	elemsCopy := make(Int16Set, len(*set), cap(*set))
	copy(elemsCopy, *set)
	return elemsCopy, nil
}

// ToSlice - Returns a slice of native datatype from the set
func (set *Int16Set) ToSlice() ([]int16, error) {
	if set.IsEmpty() {
		return nil, errors.New(common.EmptySet)
	}

	result := make([]int16, len(*set))
	for i, v := range *set {
		result[i] = int16(v)
	}
	return result, nil
}
