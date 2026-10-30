package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arg1 := os.Args[1]

	if arg1[0] != '[' || arg1[len(arg1)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	arg1 = arg1[1 : len(arg1)-1]

	slice1 := []string{}
	elem := ""
	for _, r := range arg1 {
		if r == ' ' {
			continue
		}
		if r == ',' {
			slice1 = append(slice1, elem)
			elem = ""
			continue
		}
		elem = elem + string(r)
	}
	if len(elem) > 0 {
		slice1 = append(slice1, elem)
	}

	numbers := []int{}
	for _, el := range slice1 {
		num, ok := strconv(el)
		if ok == false {
			fmt.Println("Invalid number:", el)
			return
		}
		numbers = append(numbers, num)
	}

	arg2 := os.Args[2]
	target, ok := strconv(arg2)
	if ok == false {
		fmt.Println("Invalid target sum.")
		return
	}

	// ----- valid! ------

	// slice -> []type{}
	// array -> [n]type{}

	result := [][2]int{}
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				result = append(result, [2]int{i, j})
			}
		}
	}

	if len(result) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	fmt.Printf("Pairs with sum %d: [", target)
	for i, pair := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("[%d %d]", pair[0], pair[1])
	}
	fmt.Println("]")

}

// s = "123"
// result = 0

// r = '1' -> result = 0 -> result = 1
// r = '2' -> result = 10 -> result = 12
// r = '3' -> result = 120 -> result = 123

func strconv(s string) (int, bool) {
	result := 0
	isNegative := false

	for i, r := range s {
		if i == 0 && r == '-' {
			isNegative = true
		} else if r >= '0' && r <= '9' {
			result = result * 10
			result = result + int(r-'0')
		} else {
			return 0, false
		}
	}

	if isNegative {
		return -result, true
	}
	return result, true
}
