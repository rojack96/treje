package mapSet

import (
	"errors"
	"github.com/rojack96/treje/common"
	s "github.com/rojack96/treje/set"
	"sort"
)

type Int32Set map[int32]void

// NewInt32Set - Create a new empty set or from a slice
func NewInt32Set(elems ...int32) (Int32Set, error) {
	set := Int32Set{}

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
func (set *Int32Set) Add(elem int32) {
	(*set)[elem] = void{}
}

// Remove - Remove a specific element from a set, if the element not exists raise an error
func (set *Int32Set) Remove(elem int32) error {
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
func (set *Int32Set) Discard(elem int32) {
	delete(*set, elem)
}

/*
	Set operation methods
*/

// Union - Merges the current set with another set, but returns an error
// if there are any duplicates in the union.
func (set *Int32Set) Union(b Int32Set) (Int32Set, error) {

	for elemB := range b {
		if set.Has(elemB) {
			return nil, errors.New(common.HasDuplicates)
		}
		set.Add(elemB)
	}
	return *set, nil
}

// Intersect - Returns the elements that are present in both input sets.
func (set *Int32Set) Intersect(b Int32Set) Int32Set {
	var result Int32Set

	for k := range *set {
		if _, ok := b[k]; ok {
			result[k] = void{}
		}
	}

	return result
}

// Difference - Returns the elements that are present in the first set
// but not in the second set.
func (set *Int32Set) Difference(b Int32Set) Int32Set {
	var result Int32Set

	for k := range *set {
		if _, ok := b[k]; !ok {
			result[k] = void{}
		}
	}

	return result
}

// SymmetricDifference - Returns a new set with elements that are present in either of the two sets but not in both.
func (set *Int32Set) SymmetricDifference(b Int32Set) Int32Set {
	var (
		diff1, diff2 Int32Set
	)
	diff1 = set.Difference(b)
	diff2 = (&b).Difference(*set)

	for key := range diff1 {
		diff2[key] = void{}
	}

	return diff1
}

// IsSubsetOf - Returns true if the current set is a subset of the given set b.
func (set *Int32Set) IsSubsetOf(b Int32Set) bool {
	for key := range *set {
		if _, found := b[key]; !found {
			return false
		}
	}
	return true
}

// Equals - Returns true if the current set and set b contain the same elements.
func (set *Int32Set) Equals(b Int32Set) bool {
	return set.IsSubsetOf(b) && (&b).IsSubsetOf(*set)
}

/*
	Utility methods
*/

// Has - Return true if the element is in set, otherwise false
func (set *Int32Set) Has(elem int32) bool {
	_, ok := (*set)[elem]
	return ok
}

// IsEmpty - Return true if the set is empty, else false
func (set *Int32Set) IsEmpty() bool {
	return len(*set) == 0
}

// Clear - Remove all elements
func (set *Int32Set) Clear() {
	*set = Int32Set{}
}

// Min - Return minimum element from the set
func (set *Int32Set) Min() (int32, error) {
	if set.IsEmpty() {
		return 0, errors.New(common.EmptySet)
	}

	var (
		slice []int32
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
func (set *Int32Set) Max() (int32, error) {
	if set.IsEmpty() {
		return 0, errors.New(common.EmptySet)
	}

	var (
		slice []int32
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
func (set *Int32Set) Sum() int {
	var (
		keys []int32
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
func (set *Int32Set) Sort() error {
	var (
		originalMap Int32Set
		keys        []int32
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
func (set *Int32Set) ReverseSort() error {
	var (
		originalMap Int32Set
		keys        []int32
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

func (set *Int32Set) Copy() (Int32Set, error) {
	if set.IsEmpty() {
		return nil, errors.New(common.CopyEmpty)
	}

	elemsCopy := make(Int32Set, len(*set))

	for key := range *set {
		elemsCopy[key] = void{}
	}

	return elemsCopy, nil
}

// ToSlice - Returns a slice of native datatype from the map set
func (set *Int32Set) ToSlice() ([]int32, error) {
	if set.IsEmpty() {
		return nil, errors.New(common.EmptySet)
	}

	result := make([]int32, len(*set))
	for k := range *set {
		result = append(result, k)
	}
	return result, nil
}

// ToSet - Returns a Set entities
func (set *Int32Set) ToSet() (s.Int32Set, error) {
	var (
		slice  []int32
		result s.Int32Set
		err    error
	)

	if slice, err = set.ToSlice(); err != nil {
		return nil, err
	}

	result, _ = s.NewInt32Set(slice...)
	return result, nil
}
