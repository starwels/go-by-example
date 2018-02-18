package main

import "os"

func main() {

	panic("a problem")
	panic("panic 2")
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
