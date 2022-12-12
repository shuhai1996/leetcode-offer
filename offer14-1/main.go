package main

import "fmt"

func main()  {
	//贪心算法
	fmt.Println(cuttingRope(1000))
}

func cuttingRope(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 2
	default:
		ret :=1
		for n>4 {
			ret = ret*3%1000000007
			n-=3
		}
		return ret*n%1000000007
	}
}
