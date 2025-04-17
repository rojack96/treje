package set

import (
	"errors"
	"sort"
	"strconv"
)

type Float64 float64
type Float64Set []Float64

// NewFloat64Set - Create a new empty set or from a slice
func NewFloat64Set(elems ...float64) (Float64Set, error) {
	set := Float64Set{}

	if len(elems) == 0 {
		return set, nil
	}

	if len(elems) == 1 {
		return append(set, Float64(elems[0])), nil
	}

	elemsCopy := make([]float64, len(elems))
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
		set = append(set, Float64(n))
	}

	return set, nil
}

/*
	Manipulation set methods
*/

// Add - Append a new element to the set if and only if it is not already present
func (set *Float64Set) Add(elem Float64) error {
	if set.Has(elem) {
		return errors.New(strconv.Itoa(int(elem)) + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return nil
}

// Remove - Remove a specific element from set, if element not exist raise an error
func (set *Float64Set) Remove(elem Float64) error {
	if set.IsEmpty() {
		return errors.New(EmptySet)
	}

	originalLen := len(*set)
	set.Discard(elem)
	if len(*set) == originalLen {
		return errors.New(ElemNotExist)
	}
	return nil
}

// Discard - Remove a specific element from set
func (set *Float64Set) Discard(elem Float64) {
	result := *set
	for i, n := range result {
		if n == elem {
			*set = append(result[:i], result[i+1:]...)
			break
		}
	}
}

// Pop - Remove and return element from set at a given index (or last if none provided)
func (set *Float64Set) Pop(index ...int) (float64, error) {
	if set.IsEmpty() {
		return 0, errors.New(EmptySet)
	}

	i := len(*set) - 1
	if len(index) > 0 {
		i = index[0]
		if i < 0 || i >= len(*set) {
			return 0, errors.New(IndexOutOfRange)
		}
	}

	elem := (*set)[i]
	*set = append((*set)[:i], (*set)[i+1:]...)
	return float64(elem), nil
}

/*
	Set operation methods
*/

// Union - Merges the current set with another set, but returns an error
// if there are any duplicates in the union.
func (set *Float64Set) Union(b Float64Set) (Float64Set, error) {
	if (&b).IsEmpty() {
		return nil, errors.New(UnionEmpty)
	}

	for _, elemB := range b {
		if set.Has(elemB) {
			return *set, errors.New(HasDuplicates)
		}
		*set = append(*set, elemB)
	}
	return *set, nil
}

// Intersect - Returns the elements that are present in both input sets.
func (set *Float64Set) Intersect(b Float64Set) (Float64Set, error) {
	if (&b).IsEmpty() {
		return nil, errors.New(IntersectEmpty)
	}
	set.Sort()
	b.Sort()

	var result Float64Set
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
func (set *Float64Set) Difference(b Float64Set) (Float64Set, error) {
	var result Float64Set

	if (&b).IsEmpty() {
		return nil, errors.New(DifferenceEmpty)
	}

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
func (set *Float64Set) SymmetricDifference(b Float64Set) (Float64Set, error) {
	var (
		diff1, diff2 Float64Set
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
func (set *Float64Set) IsSubsetOf(b Float64Set) bool {
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
func (set *Float64Set) Equals(b Float64Set) bool {
	return set.IsSubsetOf(b) && (&b).IsSubsetOf(*set)
}

/*
	Utility methods
*/

// Has - Return true if element is in set, otherwise false
func (set *Float64Set) Has(elem Float64) bool {
	for _, n := range *set {
		if n == elem {
			return true
		}
	}
	return false
}

// IsEmpty - Return true if set is empty, else false
func (set *Float64Set) IsEmpty() bool {
	return len(*set) == 0
}

// Clear - Remove all element
func (set *Float64Set) Clear() {
	*set = Float64Set{}
}

// Min - Return minimum element from set
func (set *Float64Set) Min() float64 {
	if set.IsEmpty() {
		return 0
	}

	minimum := *set
	minimum.Sort()

	res := minimum[0]
	return float64(res)
}

// Max - Return maximum element from set
func (set *Float64Set) Max() float64 {
	if set.IsEmpty() {
		return 0
	}

	maximum := *set
	maximum.Sort()

	res := maximum[len(maximum)-1]
	return float64(res)
}

// Sum - Return a sum of all elements
func (set *Float64Set) Sum() int {
	total := 0

	if len(*set) > 0 {
		for _, v := range *set {
			total += int(v)
		}
	}

	return total
}

// Sort - Sort element in ascending mode
func (set *Float64Set) Sort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] < (*set)[j]
	})
}

// ReverseSort - Sort element in descending mode
func (set *Float64Set) ReverseSort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] > (*set)[j]
	})
}

/*
	Methods to manipulate set object
*/

func (set *Float64Set) Copy() (Float64Set, error) {
	if set.IsEmpty() {
		return nil, errors.New(CopyEmpty)
	}
	elemsCopy := make(Float64Set, len(*set), cap(*set))
	copy(elemsCopy, *set)
	return elemsCopy, nil
}

// ToSlice - Returns a slice of native datatype from the set
func (set *Float64Set) ToSlice() ([]float64, error) {
	if set.IsEmpty() {
		return nil, errors.New(EmptySet)
	}

	result := make([]float64, len(*set))
	for i, v := range *set {
		result[i] = float64(v)
	}
	return result, nil
}
