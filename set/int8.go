package set

import (
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Int8 int8
type Int8Set []Int8

// NewInt8Set - Create a new empty set or from a slice
func NewInt8Set(elems ...int8) Int8Set {
	set := Int8Set{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		return append(set, Int8(elems[0]))
	}

	elemsCopy := make([]int8, len(elems))
	copy(elemsCopy, elems)

	sort.Slice(elems, func(i, j int) bool {
		return elems[i] < elems[j]
	})

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

// Add - Append a new element to the set if and only if it is not already present
func (set *Int8Set) Add(elem Int8) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

// Remove - Remove a specific element from set
func (set *Int8Set) Remove(elem Int8) {
	result := *set
	for i, n := range result {
		if n == elem {
			*set = append(result[:i], result[i+1:]...)
			return
		}
	}

	panic(ElemNotExist)
}

func (set *Int8Set) Discard(elem Int8) Int8Set {
	result := *set
	for i, n := range result {
		if n == elem {
			*set = append(result[:i], result[i+1:]...)
		}
	}

	return *set
}

func (set *Int8Set) Pop() Int8Set {
	if len(*set) == 0 {
		return *set
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(len(*set))

	*set = append((*set)[:i], (*set)[i+1:]...)
	return *set
}

func (set *Int8Set) Union(elems Int8Set) Int8Set {
	*set = append(*set, elems...)
	return *set
}

func (set *Int8Set) Intersect(elems Int8Set) Int8Set {
	set.Sort()
	elems.Sort()

	var result Int8Set
	i, j := 0, 0

	for i < len(*set) && j < len(elems) {
		if (*set)[i] == elems[j] {
			if len(result) == 0 || result[len(result)-1] != (*set)[i] {
				result = append(result, (*set)[i])
			}
			i++
			j++
		} else if (*set)[i] < elems[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

func (set *Int8Set) Difference(elems Int8Set) Int8Set {
	var result Int8Set

	for _, elemA := range *set {
		found := false
		for _, elemB := range elems {
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

func (set *Int8Set) Clear() {
	*set = Int8Set{}
}

func (set *Int8Set) Min() int {
	minimum := *set
	minimum.Sort()

	res := minimum[0]
	return int(res)
}

func (set *Int8Set) Max() int {
	maximum := *set
	maximum.Sort()

	res := maximum[len(maximum)-1]
	return int(res)
}

func (set *Int8Set) Sum() int {
	total := 0
	for _, v := range *set {
		total += int(v)
	}
	return total
}

func (set *Int8Set) Sort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] < (*set)[j]
	})
}

func (set *Int8Set) ToSlice() []int8 {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]int8, len(*set))
	for i, v := range *set {
		result[i] = int8(v)
	}
	return result
}
