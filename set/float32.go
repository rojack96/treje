package set

import (
	"errors"
	"github.com/rojack96/treje/common"
	"sort"
	"strconv"
)

type (
	Float32    float32
	Float32Set []Float32
)

// NewFloat32Set - Create a new empty set or from a slice
func NewFloat32Set(elems ...float32) (Float32Set, error) {
	set := Float32Set{}

	if len(elems) == 0 {
		return set, nil
	}

	if len(elems) == 1 {
		return append(set, Float32(elems[0])), nil
	}

	elemsCopy := make([]float32, len(elems))
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
		set = append(set, Float32(n))
	}

	return set, nil
}

/*
	Manipulation set methods
*/

// Add - Append a new element to the set if and only if it is not already present
func (set *Float32Set) Add(elem Float32) error {
	if set.Has(elem) {
		return errors.New(strconv.Itoa(int(elem)) + " " + common.AlreadyExists)
	}

	*set = append(*set, elem)
	return nil
}

// Remove - Remove a specific element from a set, if the element not exists raise an error
func (set *Float32Set) Remove(elem Float32) error {
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
func (set *Float32Set) Discard(elem Float32) {
	result := *set
	for i, n := range result {
		if n == elem {
			*set = append(result[:i], result[i+1:]...)
			break
		}
	}
}

// Pop - Remove and return element from a set at a given index (or last if none provided)
func (set *Float32Set) Pop(index ...int) (float32, error) {
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
	return float32(elem), nil
}

/*
	Set operation methods
*/

// Union - Merges the current set with another set, but returns an error
// if there are any duplicates in the union.
func (set *Float32Set) Union(b Float32Set) (Float32Set, error) {
	for _, elemB := range b {
		if set.Has(elemB) {
			return *set, errors.New(common.HasDuplicates)
		}
		*set = append(*set, elemB)
	}
	return *set, nil
}

// Intersect - Returns the elements that are present in both input sets.
func (set *Float32Set) Intersect(b Float32Set) (Float32Set, error) {
	set.Sort()
	b.Sort()

	var result Float32Set
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
func (set *Float32Set) Difference(b Float32Set) (Float32Set, error) {
	var result Float32Set

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
func (set *Float32Set) SymmetricDifference(b Float32Set) (Float32Set, error) {
	var (
		diff1, diff2 Float32Set
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
func (set *Float32Set) IsSubsetOf(b Float32Set) bool {
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
func (set *Float32Set) Equals(b Float32Set) bool {
	return set.IsSubsetOf(b) && (&b).IsSubsetOf(*set)
}

/*
	Utility methods
*/

// Has - Return true if the element is in set, otherwise false
func (set *Float32Set) Has(elem Float32) bool {
	for _, n := range *set {
		if n == elem {
			return true
		}
	}
	return false
}

// IsEmpty - Return true if the set is empty, else false
func (set *Float32Set) IsEmpty() bool {
	return len(*set) == 0
}

// Clear - Remove all elements
func (set *Float32Set) Clear() {
	*set = Float32Set{}
}

// Min - Return minimum element from the set
func (set *Float32Set) Min() float32 {
	if set.IsEmpty() {
		return 0
	}

	minimum := *set
	minimum.Sort()

	res := minimum[0]
	return float32(res)
}

// Max - Return maximum element from the set
func (set *Float32Set) Max() float32 {
	if set.IsEmpty() {
		return 0
	}

	maximum := *set
	maximum.Sort()

	res := maximum[len(maximum)-1]
	return float32(res)
}

// Sum - Return a sum of all elements
func (set *Float32Set) Sum() int {
	total := 0

	if len(*set) > 0 {
		for _, v := range *set {
			total += int(v)
		}
	}

	return total
}

// Sort - Sort element in ascending mode
func (set *Float32Set) Sort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] < (*set)[j]
	})
}

// ReverseSort - Sort element in descending mode
func (set *Float32Set) ReverseSort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] > (*set)[j]
	})
}

/*
	Methods to manipulate a set object
*/

func (set *Float32Set) Copy() (Float32Set, error) {
	if set.IsEmpty() {
		return nil, errors.New(common.CopyEmpty)
	}
	elemsCopy := make(Float32Set, len(*set), cap(*set))
	copy(elemsCopy, *set)
	return elemsCopy, nil
}

// ToSlice - Returns a slice of native datatype from the set
func (set *Float32Set) ToSlice() ([]float32, error) {
	if set.IsEmpty() {
		return nil, errors.New(common.EmptySet)
	}

	result := make([]float32, len(*set))
	for i, v := range *set {
		result[i] = float32(v)
	}
	return result, nil
}
