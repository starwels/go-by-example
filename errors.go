package main

import(
	"fmt"
	"errors"
)

func F1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg int
	prob string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func F2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
        if r, e := F1(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

	for _, i := range []int{7, 42} {
		if r, e := F2(i); e != nil {
			fmt.Println("F2 failed:", e)
		} else {
			fmt.Println("F2 worked:", r)
		}
	}

	_, e := F2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
	}
	
	_, f := F2(7)
    if ae, ok := f.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
