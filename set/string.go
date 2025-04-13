package set

type String string

type StringSet []String

func (set *StringSet) Add(elem String) StringSet {
	for _, w := range *set {
		if w != elem {
			continue
		}
		panic(elem + " " + AlreadyExists)
	}

	*set = append(*set, elem)
	return *set
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
