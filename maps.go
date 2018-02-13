package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	fmt.Println("Deleteing 'k2'")
	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	_, prs2 := m["k1"]
	fmt.Println("prs2:", prs2)

	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map:", n)

	o := make(map[string]string)
	o["foo"] = "bar"
	o["bar"] = "baz"
	fmt.Println("My map:", o)
}
