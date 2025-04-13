package set

import (
	"math/rand"
	"strconv"
	"time"
)

type Int8 int8
type Int8Set []Int8

func (set *Int8Set) Add(elem Int8) {
	for _, n := range *set {
		if (n ^ elem) == 0 {
			panic(strconv.Itoa(int(elem)) + " " + AlreadyExists)
		}
	}

	*set = append(*set, elem)
}

func (set *Int8Set) Remove(elem Int8) {
	result := *set
	for i, n := range result {
		if (n ^ elem) == 0 {
			*set = append(result[:i], result[i+1:]...)
			return
		}
	}

	panic(ElemNotExist)
}

func (set *Int8Set) Discard(elem Int8) Int8Set {
	result := *set
	for i, n := range result {
		if (n ^ elem) == 0 {
			*set = append(result[:i], result[i+1:]...)
		}
	}

	return *set
}

func (set *Int8Set) Pop() Int8Set {
	result := *set
	if len(result) == 0 {
		return result
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(len(result))

	result = append(result[:i], result[i+1:]...)
	return result
}

func (set *Int8Set) Union(elems Int8Set) Int8Set {
	*set = append(*set, elems...)
	return *set
}

func (set *Int8Set) Intersect(elems Int8Set) Int8Set {
	st := *set
	st.Sort()
	elems.Sort()

	var result Int8Set
	i, j := 0, 0

	for i < len(st) && j < len(elems) {
		if st[i] == elems[j] {
			// evita duplicati nel risultato
			if len(result) == 0 || result[len(result)-1] != st[i] {
				result = append(result, st[i])
			}
			i++
			j++
		} else if st[i] < elems[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

func (set *Int8Set) Difference(elems Int8Set) Int8Set {
	var result Int8Set
	st := *set

	for _, elemA := range st {
		found := false
		for _, elemB := range elems {
			if elemA == elemB { // TODO usare XOR
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
	set.quickSort(*set, 0, len(*set)-1)
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

func (set *Int8Set) partition(slice []Int8, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func (set *Int8Set) quickSort(slice []Int8, low, high int) {
	if low < high {
		pivot := set.partition(slice, low, high)
		set.quickSort(slice, low, pivot-1)
		set.quickSort(slice, pivot+1, high)
	}
}
