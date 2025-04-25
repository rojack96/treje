package types

import (
	"errors"
	"github.com/rojack96/treje/common"
	s "github.com/rojack96/treje/set"
	stype "github.com/rojack96/treje/set/types"
	"sort"
	"strings"
)

type StringSet map[string]void

// String - Create a new empty set or from a slice
func (m MapSet) String(elems ...string) (StringSet, error) {
	set := StringSet{}

	if len(elems) == 0 {
		return set, nil
	}

	for _, e := range elems {
		set.Add(e)
	}
	return set, nil
}

/*
	Manipulation set methods
*/

// Add - Append a new element to the set if and only if it is not already present
func (set *StringSet) Add(elem string) {
	(*set)[elem] = void{}
}

// Remove - Remove a specific element from a set, if the element not exists raise an error
func (set *StringSet) Remove(elem string) error {
	if set.IsEmpty() {
		return errors.New(common.EmptySet)
	}

	if !set.Has(elem) {
		return errors.New(common.ElemNotExist)
	}
	set.Discard(elem)
	return nil
}

// Discard - Remove a specific element from set
func (set *StringSet) Discard(elem string) {
	delete(*set, elem)
}

/*
	Set operation methods
*/

// Union - Merges the current set with another set, but returns an error
// if there are any duplicates in the union.
func (set *StringSet) Union(b StringSet) (StringSet, error) {

	for elemB := range b {
		if set.Has(elemB) {
			return nil, errors.New(common.HasDuplicates)
		}
		set.Add(elemB)
	}
	return *set, nil
}

// Intersect - Returns the elements that are present in both input sets.
func (set *StringSet) Intersect(b StringSet) StringSet {
	result := make(StringSet)

	for k := range *set {
		if _, ok := b[k]; ok {
			result[k] = void{}
		}
	}

	return result
}

// Difference - Returns the elements that are present in the first set
// but not in the second set.
func (set *StringSet) Difference(b StringSet) StringSet {
	result := make(StringSet)

	for k := range *set {
		if _, ok := b[k]; !ok {
			result[k] = void{}
		}
	}

	return result
}

// SymmetricDifference - Returns a new set with elements that are present in either of the two sets but not in both.
func (set *StringSet) SymmetricDifference(b StringSet) StringSet {
	var (
		diff1, diff2 StringSet
	)
	diff1 = set.Difference(b)
	diff2 = (&b).Difference(*set)

	for key := range diff1 {
		diff2[key] = void{}
	}

	return diff1
}

// IsSubsetOf - Returns true if the current set is a subset of the given set b.
func (set *StringSet) IsSubsetOf(b StringSet) bool {
	for key := range *set {
		if _, found := b[key]; !found {
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
func (set *StringSet) Has(elem string) bool {
	_, ok := (*set)[elem]
	return ok
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
func (set *StringSet) Min() (string, error) {
	if set.IsEmpty() {
		return "", errors.New(common.EmptySet)
	}

	var (
		slice []string
		err   error
	)

	maximum := *set
	if err = maximum.Sort(); err != nil {
		return "", err
	}

	if slice, err = maximum.ToSlice(); err != nil {
		return "", err
	}
	return slice[0], nil
}

// Max - Return maximum element from the set
func (set *StringSet) Max() (string, error) {
	if set.IsEmpty() {
		return "", errors.New(common.EmptySet)
	}

	var (
		slice []string
		err   error
	)

	maximum := *set
	if err = maximum.Sort(); err != nil {
		return "", err
	}

	if slice, err = maximum.ToSlice(); err != nil {
		return "", err
	}

	return slice[len(slice)-1], nil
}

// Concat - Return a string concat of all elements with a separator
func (set *StringSet) Concat(separator string) string {
	var (
		keys []string
		err  error
	)

	if keys, err = set.ToSlice(); err != nil {
		return ""
	}

	return strings.Join(keys, separator)
}

// Sort - Sort element in ascending mode
func (set *StringSet) Sort() error {
	var (
		originalMap StringSet
		keys        []string
		err         error
	)
	if originalMap, err = set.Copy(); err != nil {
		return err
	}

	if keys, err = set.ToSlice(); err != nil {
		return err
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, key := range keys {
		(*set)[key] = originalMap[key]
	}

	return nil
}

// ReverseSort - Sort element in descending mode
func (set *StringSet) ReverseSort() error {
	var (
		originalMap StringSet
		keys        []string
		err         error
	)
	if originalMap, err = set.Copy(); err != nil {
		return err
	}

	if keys, err = set.ToSlice(); err != nil {
		return err
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	for _, key := range keys {
		(*set)[key] = originalMap[key]
	}

	return nil
}

/*
	Methods to manipulate a set object
*/

func (set *StringSet) Copy() (StringSet, error) {
	if set.IsEmpty() {
		return nil, errors.New(common.CopyEmpty)
	}

	elemsCopy := make(StringSet, len(*set))

	for key := range *set {
		elemsCopy[key] = void{}
	}

	return elemsCopy, nil
}

// ToSlice - Returns a slice of native datatype from the map set
func (set *StringSet) ToSlice() ([]string, error) {
	if set.IsEmpty() {
		return nil, errors.New(common.EmptySet)
	}

	result := make([]string, 0)
	for k := range *set {
		result = append(result, k)
	}
	return result, nil
}

// ToSet - Returns a Set entities
func (set *StringSet) ToSet() (stype.StringSet, error) {
	var (
		slice  []string
		result stype.StringSet
		err    error
	)

	if slice, err = set.ToSlice(); err != nil {
		return nil, err
	}

	result, _ = s.New().String(slice...)
	return result, nil
}
