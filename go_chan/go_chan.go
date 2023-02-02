package main

import (
	"fmt"
)

func main() {
	ch := make(chan Person)
	// go func() {
	// 	ch <- "hi" // send
	// }()
	// msg := <-ch // receive
	// fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			// msg := fmt.Sprintf("message #%d", i+1)
			person := Person{
				Age: i + 1,
			}
			ch <- person
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got:", msg)
	}
	msg := <-ch
	fmt.Println(msg)

	msg, ok := <-ch // ch is closed
	fmt.Printf("closed: %#v (ok=%v)", msg, ok)

}

type Person struct {
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("Age: %d", p.Age)
}
