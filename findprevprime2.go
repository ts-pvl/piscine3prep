package piscine

func FindPrevPrime2(nb int) int {
	if nb < 2 {
		return 0
	}

	numbers := make([]bool, nb+1)

	// заполнение
	for i := 2; i <= nb; i++ {
		numbers[i] = true
	}

	for i := 2; i <= nb; i++ {
		if numbers[i] == false {
			continue
		}
		for j := i + i; j <= nb; j += i {
			numbers[j] = false
		}
	}
	for i := nb; i >= 2; i-- {
		if numbers[i] == true {
			return i
		}
	}
	return 0

}
