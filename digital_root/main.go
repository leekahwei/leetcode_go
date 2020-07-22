package main

import "fmt"

func digital_root(x int) int {
	return (x-1)%9 + 1
}

func main() {
	fmt.Print(digital_root(28732))
}
