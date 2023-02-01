package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(safeDiv(1, 0))
}

func safeDiv(a, b int) (q int, err error) {
	// q & err are local variables in safeDiv()
	defer func() {
		// the problem here is that there's no guarantee that e is an error
		// e's type is any (or interface{}) *not* error
		// this is beacuse we can panic with a string panic("this is a panic")
		// Long story short - this is not a good way to program in Go.

		if e := recover(); e != nil {
			log.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	return a / b, nil
}

func div(a, b int) int {
	return a / b
}
