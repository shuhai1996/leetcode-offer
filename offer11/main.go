package main

import "fmt"

func main()  {
	fmt.Println(minArray([]int{3,4,5,6,2,3}))
}

//二分查找
func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	p := 0 //头指针
	q := len(numbers)-1 //尾指针
	mid := p //中间指针

	for p < q{
		mid = p+(q-p)/2
		if numbers[mid] < numbers[q] {
			q = mid //收缩边界
		} else if numbers[mid] > numbers[q]{
			p = mid + 1
		} else {
			q--
		}
	}

	return numbers[p]
}