package main

//每一个可独立运行的 Go 程序，必定包含一个 package main。
//在这个 main 包中必定包含一个入口函数 main，而这个函数既没有参数，也没有返回值。

import "fmt"

func main() {
	m := []int{3, 2, 4}
	for index, value := range m {
		fmt.Println(index, value)
	}
}
