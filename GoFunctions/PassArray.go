package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	foo(a)
	fmt.Print(a)
}

func foo(x [3]int) {
	x[0] = x[0] + 1
}
