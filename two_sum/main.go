package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i, value := range nums {
		count := i
		for count < len(nums)-1 {
			total := value + nums[count+1]
			if total == target {
				result = append(result, i, count+1)
				break
			}
			count++
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 5, 6}
	fmt.Println(twoSum(nums, 3))
}
