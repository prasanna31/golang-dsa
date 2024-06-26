package Array

// merge merges two sorted slices into a single sorted slice
func (arr *AnySlice) merge(left, right AnySlice, cmp Comparator) AnySlice {
	result := make(AnySlice, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if cmp(left[i], right[j]) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// mergeSort recursively sorts the slice using the merge sort algorithm
func (arr *AnySlice) mergeSort(cmp Comparator) AnySlice {
	if len(*arr) <= 1 {
		return *arr
	}

	mid := len(*arr) / 2
	leftSlice := AnySlice((*arr)[:mid])
	rightSlice := AnySlice((*arr)[mid:])

	left := leftSlice.mergeSort(cmp)
	right := rightSlice.mergeSort(cmp)

	return arr.merge(left, right, cmp)
}

// Comparator defines a function that compares two elements and returns true if a <= b
type Comparator func(a, b any) bool

// intComparator is a comparator function for int type elements
func intComparator(a, b any) bool {
	return a.(int) <= b.(int)
}
