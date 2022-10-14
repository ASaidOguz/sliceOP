package SliceOp

//DeleteOrder takes slice of any kind ,a i index and delete the index element
//and returns the slice with the same order
func DeleteOrder[T any](s []T, i int) []T {

	s = append(s[:i], s[i+1:]...)
	s = s[:i+copy(s[i:], s[i:])]

	newSlice := make([]T, 0, len(s)-1)
	copy(newSlice, s)
	newSlice = append([]T(nil), s...)

	newSlice = append(s[:0:0], s...)
	return newSlice
}

//DeleteOrderless takes slice of any kind ,a i index and delete the index element
//and returns the slice without order
func DeleteOrderless[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]

	newSlice := make([]T, 0, len(s)-1)
	copy(newSlice, s)
	newSlice = append([]T(nil), s...)

	newSlice = append(s[:0:0], s...)
	return newSlice
}

//AppendVector takes any kind of slice , combine them with first slice given
//and return the new slice,
func AppendVector[T any](s []T, b []T) []T {

	s = append(s, b...)

	newSlice := make([]T, 0, len(s)-1)
	copy(newSlice, s)
	newSlice = append([]T(nil), s...)

	newSlice = append(s[:0:0], s...)
	return newSlice

}

func Cut[T any](s []T, i int, j int) []T {
	s = append(s[:i], s[j:]...)

	newSlice := make([]T, 0, len(s)-1)
	copy(newSlice, s)
	newSlice = append([]T(nil), s...)

	newSlice = append(s[:0:0], s...)
	return newSlice
}

func Extend[T any](s []T, j int) []T {
	s = append(s, make([]T, j)...)

	newSlice := make([]T, 0, len(s)-1)
	copy(newSlice, s)
	newSlice = append([]T(nil), s...)

	newSlice = append(s[:0:0], s...)
	return newSlice

}

func Insert[T any](s []T, x T, i int) []T {
	s = append(s[:i], append([]T{x}, s[i:]...)...)

	newSlice := make([]T, 0, len(s)-1)
	copy(newSlice, s)
	newSlice = append([]T(nil), s...)

	newSlice = append(s[:0:0], s...)
	return newSlice
}

func InsertVector[T any](s []T, x []T, i int) []T {
	s = append(s[:i], append(x, s[i:]...)...)

	newSlice := make([]T, 0, len(s)-1)
	copy(newSlice, s)
	newSlice = append([]T(nil), s...)

	newSlice = append(s[:0:0], s...)
	return newSlice
}

func Push[T any](s []T, x T) []T {
	s = append(s, x)

	newSlice := make([]T, 0, len(s)-1)
	copy(newSlice, s)
	newSlice = append([]T(nil), s...)

	newSlice = append(s[:0:0], s...)
	return newSlice
}

func PushFront[T any](s []T, x T) []T {
	s = append([]T{x}, s...)

	newSlice := make([]T, 0, len(s)-1)
	copy(newSlice, s)
	newSlice = append([]T(nil), s...)

	newSlice = append(s[:0:0], s...)
	return newSlice
}
