package main

import "fmt"

func bubbleSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 冒泡排序核心实现代码
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
	}

	return nums
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	nums = bubbleSort(nums)
	fmt.Println(nums)
}
