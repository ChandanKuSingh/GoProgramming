package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "python"
}
func main() {

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is ", g.name)
}
