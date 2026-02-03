package cards

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	return slice[index]
}

func SetItem(slice []int, index, value int) []int {
	if index >= 0 && index < len(slice) {
		slice[index] = value
	} else {
		slice = append(slice, value)
	}

	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	totalLen := len(slice) + len(values)
	newSlice := make([]int, totalLen)
	copy(newSlice, values)
	copy(newSlice[len(values):], slice)
	return newSlice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
