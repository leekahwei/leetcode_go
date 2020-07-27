package main

import (
	"fmt"
	"strconv"
)

func findAndReplacePattern(words []string, pattern string) []string {
	findPattern := func(word string) string {
		foundCharMap := make(map[string]int)
		patternCode := ""
		i := 0
		difference := 0
		for i < len(word) {
			char := string([]rune(word)[i])
			if i == 0 {
				foundCharMap[char] = difference
				patternCode += strconv.Itoa(difference)
			} else {
				if _, found := foundCharMap[char]; found {
					patternCode += strconv.Itoa(foundCharMap[char])
				} else {
					if char == string([]rune(word)[i-1]) {
						patternCode += strconv.Itoa(difference)
						foundCharMap[char] = difference
					} else {
						difference += 1
						patternCode += strconv.Itoa(difference)
						foundCharMap[char] = difference
					}
				}
			}
			i++
		}
		return patternCode
	}

	resultArr := []string{}
	patternCode := findPattern(pattern)
	for i, value := range words {
		if patternCode == findPattern(value) {
			resultArr = append(resultArr, words[i])
		}
	}
	return resultArr
}

// ["abc","cba","xyx","yxx","yyx"]
// "abc"
func main() {
	try := []string{"abc", "cba", "xyx", "yxx", "yyx"}
	fmt.Println(findAndReplacePattern(try, "abc"))
}
