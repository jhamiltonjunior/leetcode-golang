package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			possibleTarget := nums[i] + nums[j]
			if possibleTarget == target {
				return []int{
					i,
					j,
				}
			}

		}
	}

	return []int{0}
}
