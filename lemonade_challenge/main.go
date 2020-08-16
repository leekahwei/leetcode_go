package main

import "fmt"

func lemonadeChange(bills []int) bool {
	ableToPay := true
	currentSaving := make(map[int]int)
	for i, bill := range bills {
		if i == 0 && bill != 5 {
			ableToPay = false
			break
		} else {
			if bill == 10 {
				if currentSaving[5] < 1 {
					ableToPay = false
					break
				} else {
					currentSaving[10]++
					currentSaving[5]--
				}
			} else if bill == 20 {
				if currentSaving[10] < 1 {
					if currentSaving[5] < 3 {
						ableToPay = false
						break
					} else {
						currentSaving[20]++
						currentSaving[5] -= 3
					}
				} else {
					if currentSaving[5] < 1 {
						ableToPay = false
						break
					} else {
						currentSaving[20]++
						currentSaving[10]--
						currentSaving[5]--
					}
				}
			} else {
				currentSaving[5]++
			}
		}
	}

	return ableToPay
}

func main() {
	// bills := []int{5, 5, 5, 10, 20}
	bills := []int{5, 5, 10}
	// bills := []int{5, 5, 10, 10, 20}

	fmt.Println(lemonadeChange(bills))
}
