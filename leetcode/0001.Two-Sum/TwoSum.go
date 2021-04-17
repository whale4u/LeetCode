package gotest

import "fmt"

// 自己写的bug
func TwoSum(nums []int, target int) []int {
	for index1, value1 := range nums {
		for index2, value2 := range nums {
			if (value1 + value2) == target {
				// 如果存在就返回数组
				return []int{index1, index2}
			}
		}
	}
	// 如果不存在就返回空数组
	return []int{}
}

// 测试用例：{"1", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
// 其它人的方法
func TwoSum1(nums []int, target int) []int {
	// 创建1个map
	m := make(map[int]int)
	// 遍历数组
	for k, v := range nums {
		// 如果k-v在m中存在
		if idx, ok := m[target-v]; ok {
			fmt.Println("index, value: ", k, v)
			fmt.Println(m)
			// 返回数组
			return []int{idx, k}
		}
		//如果不存在，把数组中的index，value倒序存入map中
		m[v] = k
		// fmt.Println(m)
	}
	return nil
}
