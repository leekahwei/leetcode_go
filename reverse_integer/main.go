package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	// check if value provided is within the 32 bit range
	// check if the value provided contain negative
	// perform an integer reverse sorting using push and pop
	checkExceedRange := func(param int) int {
		int32Range := int(math.Pow(2, 31))
		if param >= -int32Range && param <= int32Range-1 {
			return param
		}
		return 0
	}

	result := 0
	if checkExceedRange(x) != 0 {
		isNegative := false
		stringX := strconv.Itoa(x)
		if strings.Contains(stringX, "-") {
			isNegative = true
			stringX = strings.TrimLeft(stringX, "-")
		}

		runes := []rune(stringX)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result, _ = strconv.Atoi(string(runes))
		if isNegative == true {
			result = -result
		}
		result = checkExceedRange(result)
	}
	return result
}

func main() {
	// reverse(2147483648)  // test case of more than max 32 signed int
	// reverse(-2147483649) // test case of more than min 32 signed int
	fmt.Println(reverse(-123))
}
