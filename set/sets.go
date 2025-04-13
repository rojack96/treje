package set

func NewInt8Set(elems ...int8) Int8Set {
	if len(elems) == 0 {
		return Int8Set{}
	}

	set := Int8Set{}
	if len(elems) == 1 {
		e := Int8(elems[0])
		set = append(set, e)
		return set
	}

	quickSort8(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] != elems[i-1] {
			set = append(set, Int8(elems[i]))
		} else {
			panic("set has duplicates")
		}
	}

	return set
}

func NewInt16Set(elems ...int16) Int16Set {
	if len(elems) == 0 {
		return Int16Set{}
	}

	set := Int16Set{}
	if len(elems) == 1 {
		e := Int16(elems[0])
		set = append(set, e)
		return set
	}

	quickSort16(elems, 0, len(elems)-1)

	for i := 1; i < len(elems); i++ {
		if elems[i] != elems[i-1] {
			set = append(set, Int16(elems[i]))
		} else {
			panic("set has duplicates")
		}
	}

	return set
}

func NewInt32Set() Int32Set {
	return Int32Set{}
}

func NewInt64Set() Int64Set {
	return Int64Set{}
}

func NewIntSet() IntSet {
	return IntSet{}
}

/* ---------- Sort functions ---------- */

/* ---------- Int8 ---------- */
func partition8(slice []int8, low, high int) int {
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

func quickSort8(slice []int8, low, high int) {
	if low < high {

		pivot := partition8(slice, low, high)

		quickSort8(slice, low, pivot-1)
		quickSort8(slice, pivot+1, high)
	}
}

/* ---------- Int16 ---------- */

func partition16(slice []int16, low, high int) int {
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

func quickSort16(slice []int16, low, high int) {
	if low < high {

		pivot := partition16(slice, low, high)

		quickSort16(slice, low, pivot-1)
		quickSort16(slice, pivot+1, high)
	}
}

/* ---------- Int32 ---------- */

func partition32(slice []int32, low, high int) int {
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

func quickSort32(slice []int32, low, high int) {
	if low < high {

		pivot := partition32(slice, low, high)

		quickSort32(slice, low, pivot-1)
		quickSort32(slice, pivot+1, high)
	}
}

/* ---------- Int64 ---------- */

func partition64(slice []int64, low, high int) int {
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

func quickSort64(slice []int64, low, high int) {
	if low < high {

		pivot := partition64(slice, low, high)

		quickSort64(slice, low, pivot-1)
		quickSort64(slice, pivot+1, high)
	}
}
