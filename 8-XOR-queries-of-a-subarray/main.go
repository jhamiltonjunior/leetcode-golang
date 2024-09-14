package main

import (
	"fmt"
)

func main() {
	arr := []int{
		1,
		3,
		4,
		8,
	}

	queries := [][]int{
		{0, 4},
		{1, 2},
		{0, 3},
		{3, 3},
	}

	fmt.Println(xorQueries(arr, queries))
}

// ^
func xorQueries(arr []int, queries [][]int) []int {
	var arr2 []int

	for i := 0; i < len(queries); i++ {
		var poweredNum int

		for j := queries[i][0]; j <= queries[i][1]; j++ {
			if len(arr) <= j {
				break
			}

			poweredNum ^= arr[j]
		}

		arr2 = append(arr2, poweredNum)
	}

	return arr2
}
