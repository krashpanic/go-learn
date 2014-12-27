package main

import "fmt"

var (
	x  int
	px *int
)

func main() {
	x = 15
	px = &x

	fmt.Println(x, px)
}
