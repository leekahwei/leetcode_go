package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {
	arr_people := make([]int, num_people)
	given_candies := 1

	for candies > 0 {
		i := 0
		for i < num_people {
			if given_candies <= candies {
				arr_people[i] += given_candies
			} else {
				arr_people[i] += candies
			}
			i++
			candies -= given_candies
			given_candies += 1
			if candies < 0 {
				break
			}
		}
	}

	return arr_people
}

func main() {
	fmt.Println(distributeCandies(10, 3))
}
