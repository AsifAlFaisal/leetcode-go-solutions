package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	now := time.Now()
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
	fmt.Println(Median(nums1))
	fmt.Println(time.Since(now))
}

func Median(x []int) float64 {
	l := len(x)
	idx := l / 2
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		return float64(x[idx-1]+x[idx]) / 2.0
	} else {
		return float64(x[idx])
	}
}
