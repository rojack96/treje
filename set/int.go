package set

import (
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Int int

type IntSet []Int

func NewIntSet(elems ...int) IntSet {
	set := IntSet{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, Int(elems[0]))
		return set
	}

	sort.Ints(elems)

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
		set = append(set, Int(elems[i]))
	}

	for _, n := range elems {
		set = append(set, Int(n))
	}

	return set
}

func (set *IntSet) Add(elem Int) {
	for _, n := range *set {
		if n == elem {
			panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

// Remove - Remove a specific element from set
func (set *IntSet) Remove(elem Int) {
	result := *set
	for i, n := range result {
		if n == elem {
			*set = append(result[:i], result[i+1:]...)
			return
		}
	}

	panic(ElemNotExist)
}

func (set *IntSet) Discard(elem Int) IntSet {
	result := *set
	for i, n := range result {
		if n == elem {
			*set = append(result[:i], result[i+1:]...)
		}
	}

	return *set
}

func (set *IntSet) Pop() IntSet {
	if len(*set) == 0 {
		return *set
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(len(*set))

	*set = append((*set)[:i], (*set)[i+1:]...)
	return *set
}

func (set *IntSet) Union(elems IntSet) IntSet {
	*set = append(*set, elems...)
	return *set
}

func (set *IntSet) Intersect(elems IntSet) IntSet {
	set.Sort()
	elems.Sort()

	var result IntSet
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

func (set *IntSet) Difference(elems IntSet) IntSet {
	var result IntSet

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

func (set *IntSet) Clear() {
	*set = IntSet{}
}

func (set *IntSet) Min() int {
	minimum := *set
	minimum.Sort()

	res := minimum[0]
	return int(res)
}

func (set *IntSet) Max() int {
	maximum := *set
	maximum.Sort()

	res := maximum[len(maximum)-1]
	return int(res)
}

func (set *IntSet) Sum() int {
	total := 0
	for _, v := range *set {
		total += int(v)
	}
	return total
}

func (set *IntSet) Sort() {
	sort.Slice(*set, func(i, j int) bool {
		return (*set)[i] < (*set)[j]
	})
}

func (set *IntSet) ToSlice() []int {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]int, len(*set))
	for i, v := range *set {
		result[i] = int(v)
	}
	return result
}
