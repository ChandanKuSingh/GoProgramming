package main

import "fmt"

func main() {
	fmt.Print(getMaxValue(1, 2, 3, 23, 45, 9))
}

func getMaxValue(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}
