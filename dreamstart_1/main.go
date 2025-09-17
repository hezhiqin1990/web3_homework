package main

import "fmt"

func main() {
	singleNumber([]int{4, 1, 2, 1, 2})
}

func singleNumber(nums []int) {
	var coutMap = make(map[int]int)

	for _, v := range nums {
		coutMap[v]++
	}

	for k, v := range coutMap {
		if v == 1 {
			fmt.Println("次数为1的数字	是:", k)
			break
		}
	}

}
