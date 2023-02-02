package main

import (
	"fmt"
	"time"
)

/*
for every value "n" in values, spin a goroutine that will
- sleep "n" milliseconds
- send "n" over a channel

in the function body, collect values from the channel to a slice and return it.
*/

func sleepSort(values []int) []int {
	ch := make(chan int)

	for _, n := range values {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}(n)
	}

	var out []int
	for range values {
		n := <-ch
		out = append(out, n)
	}
	return nil
}

func main() {
	values := []int{4, 6, 2, 6, 7, 1}
	sorted := sleepSort(values)
	fmt.Println(sorted)

	// ch := make(chan Person)
	// // go func() {
	// // 	ch <- "hi" // send
	// // }()
	// // msg := <-ch // receive
	// // fmt.Println(msg)
	//
	// go func() {
	// 	for i := 0; i < 3; i++ {
	// 		// msg := fmt.Sprintf("message #%d", i+1)
	// 		person := Person{
	// 			Age: i + 1,
	// 		}
	// 		ch <- person
	// 	}
	// 	close(ch)
	// }()
	//
	// for msg := range ch {
	// 	fmt.Println("got:", msg)
	// }
	// msg := <-ch
	// fmt.Println(msg)
	//
	// msg, ok := <-ch // ch is closed
	// fmt.Printf("closed: %#v (ok=%v)", msg, ok)

	// ch <- "hi" // ch is closed - panic

}

type Person struct {
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("Age: %d", p.Age)
}
