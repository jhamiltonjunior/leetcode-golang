package main

import (
	"fmt"
)

func main() {
	values := []int{0, 2, 1, 5, 3, 4}

	fmt.Println(buildArray(values))
}

func buildArray(nums []int) []int {
	nums2 := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		nums2[i] = nums[nums[i]]
	}

	return nums2
}
