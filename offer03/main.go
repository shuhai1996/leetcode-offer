package main

import "fmt"

func main() {
	//利用hash解决问题
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}

func findRepeatNumber(nums []int) int {
	ret := 0
	m := make(map[int]int) //定义map
	for _, value := range nums {
		if _, ok := m[value]; ok { //判断map中是否有该元素
			ret = value
			return ret
		}
		m[value] = 1 //往map里设置索引为value，值为1的元素
	}
	return ret
}
