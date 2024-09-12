package main

import (
	"fmt"
)

func main() {
	fmt.Println(scoreOfString("zuxazobvak")) //131
}

func scoreOfString(s string) int {
	nums := map[int]int{}

	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			nums[i] = int(s[i-1] - s[i])
			continue
		}
		nums[i] = int(s[i] - s[i-1])
	}

	sum := 0

	for _, v := range nums {
		sum += v
	}

	return int(sum)
}
