package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	temp_out := map[int]int{}

	for i, v1 := range nums {
		if key, value := temp_out[target-v1]; value {
			return []int{key, i}
		}
		temp_out[v1] = i
	}
	return nil
}
