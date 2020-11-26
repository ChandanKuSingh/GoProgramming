package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	passArrayPointer(&a)
	fmt.Println(a)
}

func passArrayPointer(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}
