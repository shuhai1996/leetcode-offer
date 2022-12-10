package main

import "fmt"

func main() {
	fmt.Println(fib(2))
}

//注意不能用递归，递归中很多计算都是重复的，有性能影响
func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		var ret = 0
		nOne, nTwo := 1, 0
		for i := 2; i <= n; i++ {
			ret = (nOne + nTwo)%1000000007//取模
			nTwo = nOne
			nOne = ret
		}
		return ret
	}
}
