package Array

func InitArray() []any {
	// Initialize an array
	return make([]any, 0)
}

// Define a custom type for the slice
type AnySlice []any

// InsertMiddle inserts a value at a specified index
func (arr *AnySlice) InsertMiddle(index int, value any) {
	*arr = append((*arr)[:index], append([]any{value}, (*arr)[index:]...)...)
}

func (arr *AnySlice) InsertEnd(value any) []any {
	// Insert at the end
	return append((*arr), value)
}

func (arr *AnySlice) InsertStart(value any) []any {
	// Insert at the start
	return append([]any{value}, (*arr)...)
}

func (arr *AnySlice) DeleteByIndex(index int) []any {
	// Delete by index
	return append((*arr)[:index], (*arr)[index+1:]...)
}

func (arr *AnySlice) DeleteByValue(value any) []any {
	// Delete by value
	for i, v := range *arr {
		if v == value {
			return append((*arr)[:i], (*arr)[i+1:]...)
		}
	}
	return (*arr)
}

func (arr *AnySlice) Get(index int) any {
	// Search for a value
	return (*arr)[index]
}

func (arr *AnySlice) Display() {
	// Display the array
	for _, v := range *arr {
		print(v, " ")
	}
	println()
}

func (arr *AnySlice) DisplayReverse() {
	// Display the (*arr)ay in reverse
	for i := len((*arr)) - 1; i >= 0; i-- {
		print((*arr)[i], " ")
	}
	println()
}

func (arr *AnySlice) DisplayAtIndex(index int) {
	// Display the value at a given index
	println((*arr)[index])
}

func (arr *AnySlice) Reverse() []any {
	// Reverse the (*arr)ay
	for i := 0; i < len((*arr))/2; i++ {
		(*arr)[i], (*arr)[len((*arr))-1-i] = (*arr)[len((*arr))-1-i], (*arr)[i]
	}
	return (*arr)
}

func (arr *AnySlice) Sort(cmp Comparator) []any {
	// Sort the (*arr)ay
	return arr.mergeSort(cmp)
}
func (arr *AnySlice) Length() int {
	// Get the length of the (*arr)ay
	return len(*arr)
}
