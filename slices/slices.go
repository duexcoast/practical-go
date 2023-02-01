package main

import (
	"fmt"
	"reflect"
	"sort"
)

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

	vs := []float64{2, 1, 3}
	fmt.Println(median(vs))
	fmt.Println(vs)

	fmt.Println(median(nil))
	fmt.Println(reflect.TypeOf(2))

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

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}
	// copy in order to not change values of underlying slice
	nums := make([]float64, len(values))
	copy(nums, values)

	sort.Float64s(nums)
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i], nil
	}
	v := (nums[i-1] + nums[i]) / 2
	return v, nil
}

func concat(s1, s2 []string) []string {
	// restriction: no "for" loops
	// s1 = append(s1, s2...)
	// return s1
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}
