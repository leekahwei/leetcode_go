package main

import "fmt"

func checkPossibility(nums []int) (isDecreasing bool) {
	i := 0
	isChanging := 0
	for i < len(nums) {
		if i == len(nums)-1 { // last integer
			if nums[i] < nums[i-1] {
				isChanging++
			}
		} else {
			if nums[i] > nums[i+1] { // is Decreasing
				if nums[i+1] >= nums[i-1] {
					fmt.Println("IN HERE")
					isChanging = 2
					break
				}
				nums[i] = nums[i+1]
				isChanging++
			}
		}
		// fmt.Println(isChanging)
		i++
	}

	if isChanging <= 1 {
		isDecreasing = true
	}

	fmt.Println(nums)
	return
}

func main() {
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))    //expected false
	fmt.Println(checkPossibility([]int{-1, 4, 2, 3}))   //expected true
	fmt.Println(checkPossibility([]int{2, 3, 3, 2, 4})) //exptected true
}
