package main

import "fmt"

func main()  {
 	fmt.Println(hammingWeight(4294967293))
}

func hammingWeight(num uint32) int {
	count:=0
	for num > 0{
		count++
		num = (num-1)&num//一个整数跟它减1的结果做与运算就相当于把它最右边的1变成0
	}

	return count

}