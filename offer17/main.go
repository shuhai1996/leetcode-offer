package main

import (
	"fmt"
	"strconv"
)

func main()  {
	//参考了大佬的题解
	//全排列 、深度遍历
	fmt.Println(printNumbers(1))
	fmt.Println(printNumbers(2))
}

func printNumbers(n int) []int {
	var ret []int
	for num := 1; num < n+1; num++ {
		for b := byte('1'); b<=byte('9') ; b++ {
			str := make([]byte, num)
			str[0] = b//起始位设置为1
			ret = print(1, str, num, ret)
		}
	}

	return ret
}

func print(index int, nums []byte, num int, ret []int) []int{
	if num == index { //num = index表示当前数字，同时也是递归结束条件
		in,_ := strconv.Atoi(string(nums))
		ret = append(ret, in)
	} else {
		for b := byte('0'); b<=byte('9') ; b++ {
			nums[index] = b//相当于设置数字后一位
			ret = print(index+1, nums, num, ret)//递归调用
		}
	}

	return ret
}