package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals)) // 输出 [[1 6] [8 10] [15 18]]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// 按区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]
		if interval[0] <= last[1] { // 有重叠
			if interval[1] > last[1] {
				last[1] = interval[1] // 合并区间
			}
		} else {
			merged = append(merged, interval) // 无重叠，直接添加
		}
	}
	return merged
}
