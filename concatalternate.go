package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	if len(slice1) == 0 {
		return slice2
	}
	if len(slice2) == 0 {
		return slice1
	}

	result := []int{}

	takeSlice1 := true
	if len(slice2) > len(slice1) {
		takeSlice1 = false
	}

	i1 := 0
	i2 := 0
	for i1 < len(slice1) && i2 < len(slice2) {
		if takeSlice1 {
			result = append(result, slice1[i1])
			i1++
			takeSlice1 = false
		} else {
			result = append(result, slice2[i2])
			i2++
			takeSlice1 = true
		}
	}

	for ; i1 < len(slice1); i1++ {
		result = append(result, slice1[i1])
	}

	for ; i2 < len(slice2); i2++ {
		result = append(result, slice2[i2])
	}

	return result

}
