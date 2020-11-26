package main

import "fmt"

func main() {
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("value = ", d)
}

func divide(f1, f2 float32) (float32, error) {
	if f2 == 0.0 {
		//panic("can't provide zero as second value")	not use panic as panic terminates the program
		return 0.0, fmt.Errorf("Can't provide 2nd variable as zero")
	}
	return f1 / f2, nil
}
