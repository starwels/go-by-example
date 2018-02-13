package main

import "fmt"

func plus(a int, b int) (c int) {
	c = a + b
	return c
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println("plus:", plus(2, 3))

	res := plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)
}
