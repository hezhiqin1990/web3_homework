package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 4, 5}
	removeDuplicates(nums)
	fmt.Println("原数组是:", nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}

	}
	fmt.Println("去重后的数组是:", nums[:i+1])
	return i + 1
}
