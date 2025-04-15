package set

import "sort"

type String string
type StringSet []String

func NewStringSet(elems ...string) StringSet {
	set := StringSet{}

	if len(elems) == 0 {
		return set
	}

	if len(elems) == 1 {
		set = append(set, String(elems[0]))
		return set
	}

	sort.Slice(elems, func(i, j int) bool {
		return elems[i] < elems[j]
	})

	for i := 1; i < len(elems); i++ {
		if elems[i] == elems[i-1] {
			panic(HasDuplicates)
		}
	}

	for _, n := range elems {
		set = append(set, String(n))
	}

	return set
}

func (set *StringSet) Add(elem String) {
	for _, w := range *set {
		if w == elem {
			panic(elem + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *StringSet) ToSlice() []string {
	if len(*set) == 0 {
		panic(EmptySet)
	}

	result := make([]string, len(*set))
	for i, v := range *set {
		result[i] = string(v)
	}
	return result
}
