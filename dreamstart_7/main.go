package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target)) // 输出 [0 1]
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	//是当前需要查找的另一个数，使得 num + (target-num) == target。j, ok := m[target-num] 表示尝试在 map 中查找 target-num，如果找到了，ok 为 true，j 是该数的索引。
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}

	return nil
}
