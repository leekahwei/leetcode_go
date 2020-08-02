package main

import "fmt"

func checkPossibility(nums []int) bool {
	nonDecreasing := true
	changeCount := 0
	i := 0

	if len(nums) > 2 {
		for i < len(nums) {
			if i == len(nums)-1 {
				if nums[i] < nums[i-1] {
					nums[i] = nums[i-1]
					changeCount++
				}
			} else if i == 0 {
				if nums[i] > nums[i+1] {
					nums[i] = nums[i+1]
					changeCount++
				}
			} else {
				if nums[i+1] < nums[i] { //found decreasing
					if i+2 < len(nums) {
						if nums[i+2] < nums[i] {
							nums[i] = nums[i-1]
							changeCount++
						} else {
							nums[i+1] = nums[i]
							changeCount++
						}
					} else {
						nums[i+1] = nums[i]
						changeCount++
					}
					if nums[i+1] < nums[i] {
						nums[i+1] = nums[i]
						changeCount++
					}
				}
			}
			if changeCount > 1 {
				nonDecreasing = false
				break
			}
			i++
		}
	}

	return nonDecreasing
}

func main() {
	fmt.Println(checkPossibility([]int{4, 2, 1}))       //expected false
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))    //expected false
	fmt.Println(checkPossibility([]int{-1, 4, 2, 3}))   //expected true
	fmt.Println(checkPossibility([]int{2, 3, 3, 2, 4})) //exptected true
	fmt.Println(checkPossibility([]int{1, 2, 4, 5, 3})) //expected true
}
