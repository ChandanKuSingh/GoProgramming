package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	var y int
	var ip *int // var ip is pointer to int

	ip = &x
	y = *ip
	fmt.Println("value of ip = ", ip)
	fmt.Println("value of y = ", y)

	ptr := new(int)
	fmt.Println("Pointer new = ", ptr)

	var appleNum int

	fmt.Printf("Number of apples ??")
	fmt.Scan(&appleNum)

	fmt.Printf("appleNum : " + strconv.Itoa(appleNum))
}
