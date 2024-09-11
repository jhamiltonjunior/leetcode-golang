package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV")) // 1994

}

func romanToInt(s string) int {
	equivalentStrToInt := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sSlice := strings.Split(s, "")
	preNum := 0
	var totalNum int

	for _, romanNum := range sSlice {
		actualNum := equivalentStrToInt[romanNum]

		if preNum == 0 {
			totalNum += actualNum
			preNum = actualNum
			continue
		}
		if preNum < actualNum {
			totalNum -= preNum
			totalNum += actualNum - preNum
			preNum = actualNum
			continue
		}
		totalNum += actualNum
		preNum = actualNum
	}
	return totalNum
}
