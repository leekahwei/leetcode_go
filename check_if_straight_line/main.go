package main

import "fmt"

// Refine after understanding dy dx
func checkStraightLine(coordinates [][]int) bool {
	isStraightLine := true
	x0 := coordinates[0][0]
	x1 := coordinates[1][0]
	y0 := coordinates[0][1]
	y1 := coordinates[1][1]
	dx := x1 - x0
	dy := y1 - y0
	for _, coordinate := range coordinates {
		if dx*(coordinate[1]-y1) != dy*(coordinate[0]-x1) {
			isStraightLine = false
		}
	}
	return isStraightLine
}

// The stupid way
// func checkStraightLine(coordinates [][]int) bool {
// 	isStraightLine := false
// 	isSlope := true
// 	isVertical := true
// 	isHorizontal := true
// 	i := 1
// 	m := 0
// 	for i < len(coordinates) {
// 		if coordinates[i][0] != coordinates[i-1][0] { // x same
// 			isHorizontal = false
// 		}
// 		if coordinates[i][1] != coordinates[i-1][1] { // y same
// 			isVertical = false
// 		}

// 		y := coordinates[i][1] - coordinates[i-1][1]
// 		x := coordinates[i][0] - coordinates[i-1][0]
// 		if x != 0 && y != 0 {
// 			if i == 1 {
// 				m = y / x
// 			} else {
// 				if y/x != m {
// 					isSlope = false
// 				}
// 			}
// 		} else {
// 			isSlope = false
// 		}
// 		i++
// 	}

// 	if isSlope == true || isVertical == true || isHorizontal == true {
// 		isStraightLine = true
// 	}
// 	return isStraightLine
// }

func main() {
	// coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}} // expect true
	// coordinates := [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}} // expect true
	// coordinates := [][]int{{1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}, {6, 1}} // expect true
	// coordinates := [][]int{{2, 1}, {4, 2}, {6, 3}} // expect true
	// coordinates := [][]int{{1, -8}, {2, -3}, {1, 2}} // expect false
	coordinates := [][]int{{1, 1}, {2, 2}, {2, 0}} // expect false
	fmt.Println(checkStraightLine(coordinates))
}
