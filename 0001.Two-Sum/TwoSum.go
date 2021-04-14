package gotest

func TwoSum(nums []int, target int) []int {
	for index1, value1 := range nums {
		for index2, value2 := range nums {
			if (value1 + value2) == target {
				return []int{index1, index2}
			}
		}
	}
	// 为什么这里还要返回？？？
	return []int{}
}
