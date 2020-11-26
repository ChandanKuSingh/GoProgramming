package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	passSlice(a)
	fmt.Println(a)
}

func passSlice(x []int) {
	x[0] = x[0] + 1
}
