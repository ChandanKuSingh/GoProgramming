package Encapsulation

import "fmt"

type Cordinate struct {
	x float64
	y float64
}

func (c *Cordinate) InitCordinate(xn, yn float64) {
	c.x = xn
	c.y = yn
}

func (c *Cordinate) Scale(v float64) {
	c.x = c.x * v
	c.y = c.y * v
}

func (c *Cordinate) PrintCordinate() {
	fmt.Println(c.x, c.y)
}
