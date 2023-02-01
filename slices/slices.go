package main

import "fmt"

func main() {
	var s []int
	fmt.Println("len", len(s)) // len is nill safe
	if s == nil {
		fmt.Println("nil slice")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 %#v\n", s2)

	s3 := s2[1:4]
	s3[0] = 1
	s3 = append(s3, 1)
	fmt.Printf("s3 %#v\n", s3)
	fmt.Printf("s2 %#v\n", s2)
	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D"}))

}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) < cap(s) {
		s = s[:len(s)+1]
	} else {
		s2 := make([]int, 2*len(s)+1) // + 1 is for the case in which we have a len 0.
		copy(s2, s)
		s = s2[:len(s)+1]
	}
	s[i] = v
	return s
}

func concat(s1, s2 []string) []string {
	// restriction: no "for" loops
	s := make([]string, 0, len(s1)+len(s2))
	s = append(s, s1...)
	s = append(s, s2...)
	return s
}
