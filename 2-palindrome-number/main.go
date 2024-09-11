package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isPalindrome(245))
}

func isPalindrome(x int) bool {
	xInt := strconv.Itoa(x)
	xSlice := strings.Split(xInt, "")
	slices.Reverse(xSlice)
	xRev := strings.Join(xSlice, "")

	if xInt == xRev {
		return true
	}

	return false
}
