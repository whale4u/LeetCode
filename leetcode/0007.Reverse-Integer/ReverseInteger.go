package gotest

import (
	"math"
)

func reverse(x int) int {
	if x < int(math.Pow(2, 31)) || (x > (int(math.Pow(2, 31)) - 1)) {
		return 0
	}
	return x
}

// 一个正数怎么取得其负值？
