package main

import "fmt"

func main() {
	x := 2
	y := 3

	x, y = y, x

	fmt.Println(x, y)
}
