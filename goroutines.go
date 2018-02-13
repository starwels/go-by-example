package main

import "fmt"

func Fun(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
	Fun("direct")

	go Fun("first goroutine")
	go Fun("second goroutine")
	go Fun("third goroutine")
	go Fun("fourth goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}

