package types

import (
	"errors"
	"github.com/rojack96/treje/common"
	s "github.com/rojack96/treje/set"
	stype "github.com/rojack96/treje/set/types"
	"sort"
)

type UintSet map[uint]void

// Uint - Create a new empty set or from a slice
func (m MapSet) Uint(elems ...uint) (UintSet, error) {
	set := UintSet{}

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
func (set *UintSet) Add(elem uint) {
	(*set)[elem] = void{}
}

// Remove - Remove a specific element from a set, if the element not exists raise an error
func (set *UintSet) Remove(elem uint) error {
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
func (set *UintSet) Discard(elem uint) {
	delete(*set, elem)
}

/*
	Set operation methods
*/

// Union - Merges the current set with another set, but returns an error
// if there are any duplicates in the union.
func (set *UintSet) Union(b UintSet) (UintSet, error) {

	for elemB := range b {
		if set.Has(elemB) {
			return nil, errors.New(common.HasDuplicates)
		}
		set.Add(elemB)
	}
	return *set, nil
}

// Intersect - Returns the elements that are present in both input sets.
func (set *UintSet) Intersect(b UintSet) UintSet {
	result := make(UintSet)

	for k := range *set {
		if _, ok := b[k]; ok {
			result[k] = void{}
		}
	}

	return result
}

// Difference - Returns the elements that are present in the first set
// but not in the second set.
func (set *UintSet) Difference(b UintSet) UintSet {
	result := make(UintSet)

	for k := range *set {
		if _, ok := b[k]; !ok {
			result[k] = void{}
		}
	}

	return result
}

// SymmetricDifference - Returns a new set with elements that are present in either of the two sets but not in both.
func (set *UintSet) SymmetricDifference(b UintSet) UintSet {
	var (
		diff1, diff2 UintSet
	)
	diff1 = set.Difference(b)
	diff2 = (&b).Difference(*set)

	for key := range diff1 {
		diff2[key] = void{}
	}

	return diff1
}

// IsSubsetOf - Returns true if the current set is a subset of the given set b.
func (set *UintSet) IsSubsetOf(b UintSet) bool {
	for key := range *set {
		if _, found := b[key]; !found {
			return false
		}
	}
	return true
}

// Equals - Returns true if the current set and set b contain the same elements.
func (set *UintSet) Equals(b UintSet) bool {
	return set.IsSubsetOf(b) && (&b).IsSubsetOf(*set)
}

/*
	Utility methods
*/

// Has - Return true if the element is in set, otherwise false
func (set *UintSet) Has(elem uint) bool {
	_, ok := (*set)[elem]
	return ok
}

// IsEmpty - Return true if the set is empty, else false
func (set *UintSet) IsEmpty() bool {
	return len(*set) == 0
}

// Clear - Remove all elements
func (set *UintSet) Clear() {
	*set = UintSet{}
}

// Min - Return minimum element from the set
func (set *UintSet) Min() (uint, error) {
	if set.IsEmpty() {
		return 0, errors.New(common.EmptySet)
	}

	var (
		slice []uint
		err   error
	)

	maximum := *set
	if err = maximum.Sort(); err != nil {
		return 0, err
	}

	if slice, err = maximum.ToSlice(); err != nil {
		return 0, err
	}
	return slice[0], nil
}

// Max - Return maximum element from the set
func (set *UintSet) Max() (uint, error) {
	if set.IsEmpty() {
		return 0, errors.New(common.EmptySet)
	}

	var (
		slice []uint
		err   error
	)

	maximum := *set
	if err = maximum.Sort(); err != nil {
		return 0, err
	}

	if slice, err = maximum.ToSlice(); err != nil {
		return 0, err
	}

	return slice[len(slice)-1], nil
}

// Sum - Return a sum of all elements
func (set *UintSet) Sum() int {
	var (
		keys []uint
		err  error
	)

	if keys, err = set.ToSlice(); err != nil {
		return 0
	}

	total := 0

	if len(keys) > 0 {
		for _, v := range keys {
			total += int(v)
		}
	}

	return total
}

// Sort - Sort element in ascending mode
func (set *UintSet) Sort() error {
	var (
		originalMap UintSet
		keys        []uint
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
func (set *UintSet) ReverseSort() error {
	var (
		originalMap UintSet
		keys        []uint
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

func (set *UintSet) Copy() (UintSet, error) {
	if set.IsEmpty() {
		return nil, errors.New(common.CopyEmpty)
	}

	elemsCopy := make(UintSet, len(*set))

	for key := range *set {
		elemsCopy[key] = void{}
	}

	return elemsCopy, nil
}

// ToSlice - Returns a slice of native datatype from the map set
func (set *UintSet) ToSlice() ([]uint, error) {
	if set.IsEmpty() {
		return nil, errors.New(common.EmptySet)
	}

	result := make([]uint, 0)
	for k := range *set {
		result = append(result, k)
	}
	return result, nil
}

// ToSet - Returns a Set entities
func (set *UintSet) ToSet() (stype.UintSet, error) {
	var (
		slice  []uint
		result stype.UintSet
		err    error
	)

	if slice, err = set.ToSlice(); err != nil {
		return nil, err
	}

	result, _ = s.New().Uint(slice...)
	return result, nil
}
