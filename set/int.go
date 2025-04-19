package set

import (
	"errors"
	"github.com/rojack96/treje/utils"
	"sort"
	"strconv"
)

type (
	Int    int
	IntSet []Int
)

// NewIntSet - Create a new empty set or from a slice
func NewIntSet(elems ...int) (IntSet, error) {
	set := IntSet{}

	if len(elems) == 0 {
		return set, nil
	}

	if len(elems) == 1 {
		return append(set, Int(elems[0])), nil
	}

	elemsCopy := make([]int, len(elems))
	copy(elemsCopy, elems)

	sort.Slice(elems, func(i, j int) bool {
		return elems[i] < elems[j]
	})

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			return nil, errors.New(utils.HasDuplicates)
		}
	}

	for _, n := range elemsCopy {
		set = append(set, Int(n))
	}

	return set, nil
}

/*
	Manipulation set methods
*/

// Add - Append a new element to the set if and only if it is not already present
func (set *IntSet) Add(elem Int) error {
	if set.Has(elem) {
		return errors.New(strconv.Itoa(int(elem)) + " " + utils.AlreadyExists)
	}

	*set = append(*set, elem)
	return nil
}

// Remove - Remove a specific element from a set, if the element not exists raise an error
func (set *IntSet) Remove(elem Int) error {
	if set.IsEmpty() {
		return errors.New(utils.EmptySet)
	}

	originalLen := len(*set)
	set.Discard(elem)
	if len(*set) == originalLen {
		return errors.New(utils.ElemNotExist)
	}
	return nil
}

// Discard - Remove a specific element from set
func (set *IntSet) Discard(elem Int) {
	result := *set
	for i, n := range result {
		if n == elem {
			*set = append(result[:i], result[i+1:]...)
			break
		}
	}
}

// Pop - Remove and return element from a set at a given index (or last if none provided)
func (set *IntSet) Pop(index ...int) (int, error) {
	if set.IsEmpty() {
		return 0, errors.New(utils.EmptySet)
	}

	i := len(*set) - 1
	if len(index) > 0 {
		i = index[0]
		if i < 0 || i >= len(*set) {
			return 0, errors.New(utils.IndexOutOfRange)
		}
	}

	elem := (*set)[i]
	*set = append((*set)[:i], (*set)[i+1:]...)
	return int(elem), nil
}

/*
	Set operation methods
*/

// Union - Merges the current set with another set, but returns an error
// if there are any duplicates in the union.
func (set *IntSet) Union(b IntSet) (IntSet, error) {

	for _, elemB := range b {
		if set.Has(elemB) {
			return *set, errors.New(utils.HasDuplicates)
		}
		*set = append(*set, elemB)
	}
	return *set, nil
}

// Intersect - Returns the elements that are present in both input sets.
func (set *IntSet) Intersect(b IntSet) (IntSet, error) {
	set.Sort()
	b.Sort()

	var result IntSet
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
func (set *IntSet) Difference(b IntSet) (IntSet, error) {
	var result IntSet

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
func (set *IntSet) SymmetricDifference(b IntSet) (IntSet, error) {
	var (
		diff1, diff2 IntSet
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
func (set *IntSet) IsSubsetOf(b IntSet) bool {
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
func (set *IntSet) Equals(b IntSet) bool {
	return set.IsSubsetOf(b) && (&b).IsSubsetOf(*set)
}

/*
	Utility methods
*/

// Has - Return true if the element is in set, otherwise false
func (set *IntSet) Has(elem Int) bool {
	for _, n := range *set {
		if n == elem {
			return true
		}
	}
	return false
}

// IsEmpty - Return true if the set is empty, else false
func (set *IntSet) IsEmpty() bool {
	return len(*set) == 0
}

// Clear - Remove all elements
func (set *IntSet) Clear() {
	*set = IntSet{}
}

// Min - Return minimum element from the set
func (set *IntSet) Min() int {
	if set.IsEmpty() {
		return 0
	}

	minimum := *set
	minimum.Sort()

	res := minimum[0]
	return int(res)
}

// Max - Return maximum element from the set
func (set *IntSet) Max() int {
	if set.IsEmpty() {
		return 0
	}

	maximum := *set
	maximum.Sort()

	res := maximum[len(maximum)-1]
	return int(res)
}

// Sum - Return a sum of all elements
func (set *IntSet) Sum() int {
	total := 0

	if len(*set) > 0 {
		for _, v := range *set {
			total += int(v)
		}
	}

	return total
}

// Sort - Sort element in ascending mode
func (set *IntSet) Sort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] < (*set)[j]
	})
}

// ReverseSort - Sort element in descending mode
func (set *IntSet) ReverseSort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] > (*set)[j]
	})
}

/*
	Methods to manipulate a set object
*/

func (set *IntSet) Copy() (IntSet, error) {
	if set.IsEmpty() {
		return nil, errors.New(utils.CopyEmpty)
	}
	elemsCopy := make(IntSet, len(*set), cap(*set))
	copy(elemsCopy, *set)
	return elemsCopy, nil
}

// ToSlice - Returns a slice of native datatype from the set
func (set *IntSet) ToSlice() ([]int, error) {
	if set.IsEmpty() {
		return nil, errors.New(utils.EmptySet)
	}

	result := make([]int, len(*set))
	for i, v := range *set {
		result[i] = int(v)
	}
	return result, nil
}
