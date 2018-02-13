package main

import "fmt"

func val() (int, int) {
	return 1, 2
}

func main() {
	a, b := val()
	fmt.Println("a:", a, "b:", b)

	_, c := val()
	fmt.Println("c:", c)

	d, _ := val()
	fmt.Println("d:", d)
}
