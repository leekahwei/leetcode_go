package main

import "fmt"

func majorityElement(nums []int) int {
	/*
	 * create a variable called highest to store the highest occurence key
	 * loop through the given array
	 * add the current element to the map as key and occurence as value
	 * compare the current key occurence with the highest key occurence
	 * if higher, change the highest to the current key
	 */
	highest := 0
	occurencDict := map[int]int{}

	for _, value := range nums {
		if _, ok := occurencDict[value]; ok { //if found
			occurencDict[value]++
		} else {
			occurencDict[value] = 1
		}
		if _, ok := occurencDict[highest]; ok {
			if occurencDict[value] > occurencDict[highest] {
				highest = value
			}
		} else {
			highest = value
		}
	}
	return highest
}

func main() {
	fmt.Print(majorityElement([]int{2, 9, 2, 8, 8, 2}))
}
