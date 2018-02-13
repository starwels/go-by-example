package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})

	//We can ommit values only using named fiels
	fmt.Println(person{name:"Fred"})
	fmt.Println(person{name: "João"})

	fmt.Println(&person{"Ann", 40})

	s := person{"Sean", 50}
	fmt.Println(s.name, s.age)
	s.age = 200
	fmt.Println(s.name, s.age)

	sp := &s

	fmt.Println(sp.age)
	sp.age = 1000
	fmt.Println(sp.age)
}
