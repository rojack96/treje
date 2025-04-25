package types

import (
	"errors"
	"github.com/rojack96/treje/common"
	"sort"
	"strings"
)

type (
	Str       string
	StringSet []Str
)

// String - Create a new empty set or from a slice
func (s Set) String(elems ...string) (StringSet, error) {
	set := StringSet{}

	if len(elems) == 0 {
		return set, nil
	}

	if len(elems) == 1 {
		return append(set, Str(elems[0])), nil
	}

	elemsCopy := make([]string, len(elems))
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
		set = append(set, Str(n))
	}

	return set, nil
}

/*
	Manipulation set methods
*/

// Add - Append a new element to the set if and only if it is not already present
func (set *StringSet) Add(elem Str) error {
	if set.Has(elem) {
		return errors.New(string(elem) + " " + common.AlreadyExists)
	}

	*set = append(*set, elem)
	return nil
}

// Remove - Remove a specific element from a set, if the element not exists raise an error
func (set *StringSet) Remove(elem Str) error {
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
func (set *StringSet) Discard(elem Str) {
	result := *set
	for i, n := range result {
		if n == elem {
			*set = append(result[:i], result[i+1:]...)
			break
		}
	}
}

// Pop - Remove and return element from a set at a given index (or last if none provided)
func (set *StringSet) Pop(index ...int) (string, error) {
	if set.IsEmpty() {
		return "", errors.New(common.EmptySet)
	}

	i := len(*set) - 1
	if len(index) > 0 {
		i = index[0]
		if i < 0 || i >= len(*set) {
			return "", errors.New(common.IndexOutOfRange)
		}
	}

	elem := (*set)[i]
	*set = append((*set)[:i], (*set)[i+1:]...)
	return string(elem), nil
}

/*
	Set operation methods
*/

// Union - Merges the current set with another set, but returns an error
// if there are any duplicates in the union.
func (set *StringSet) Union(b StringSet) (StringSet, error) {

	for _, elemB := range b {
		if set.Has(elemB) {
			return *set, errors.New(common.HasDuplicates)
		}
		*set = append(*set, elemB)
	}
	return *set, nil
}

// Intersect - Returns the elements that are present in both input sets.
func (set *StringSet) Intersect(b StringSet) (StringSet, error) {
	set.Sort()
	b.Sort()

	var result StringSet
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
func (set *StringSet) Difference(b StringSet) (StringSet, error) {
	var result StringSet

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
func (set *StringSet) SymmetricDifference(b StringSet) (StringSet, error) {
	var (
		diff1, diff2 StringSet
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
func (set *StringSet) IsSubsetOf(b StringSet) bool {
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
func (set *StringSet) Equals(b StringSet) bool {
	return set.IsSubsetOf(b) && (&b).IsSubsetOf(*set)
}

/*
	Utility methods
*/

// Has - Return true if the element is in set, otherwise false
func (set *StringSet) Has(elem Str) bool {
	for _, n := range *set {
		if n == elem {
			return true
		}
	}
	return false
}

// IsEmpty - Return true if the set is empty, else false
func (set *StringSet) IsEmpty() bool {
	return len(*set) == 0
}

// Clear - Remove all elements
func (set *StringSet) Clear() {
	*set = StringSet{}
}

// Min - Return minimum element from the set
func (set *StringSet) Min() string {
	if set.IsEmpty() {
		return ""
	}

	minimum := *set
	minimum.Sort()

	res := minimum[0]
	return string(res)
}

// Max - Return maximum element from the set
func (set *StringSet) Max() string {
	if set.IsEmpty() {
		return ""
	}

	maximum := *set
	maximum.Sort()

	res := maximum[len(maximum)-1]
	return string(res)
}

// Concat - Return a string concat of all elements with a separator
func (set *StringSet) Concat(separator string) string {
	result, err := set.ToSlice()
	if err != nil {
		return ""
	}

	return strings.Join(result, separator)
}

// Sort - Sort element in ascending mode
func (set *StringSet) Sort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] < (*set)[j]
	})
}

// ReverseSort - Sort element in descending mode
func (set *StringSet) ReverseSort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] > (*set)[j]
	})
}

/*
	Methods to manipulate a set object
*/

func (set *StringSet) Copy() (StringSet, error) {
	if set.IsEmpty() {
		return nil, errors.New(common.CopyEmpty)
	}
	elemsCopy := make(StringSet, len(*set), cap(*set))
	copy(elemsCopy, *set)
	return elemsCopy, nil
}

// ToSlice - Returns a slice of native datatype from the set
func (set *StringSet) ToSlice() ([]string, error) {
	if set.IsEmpty() {
		return nil, errors.New(common.EmptySet)
	}

	result := make([]string, len(*set))
	for i, v := range *set {
		result[i] = string(v)
	}
	return result, nil
}
