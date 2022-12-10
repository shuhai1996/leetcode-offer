package main

import "fmt"

func main()  {
	fmt.Println(numWays(7))
}

func numWays(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 1
	case 2:
		return 2
	default:
		ret, numOne, numTwo := 0, 1, 0
		for i:=1; i<=n ;i++ {
			ret = (numOne + numTwo)%1000000007
			numTwo = numOne
			numOne = ret
		}
		return ret
	}
}
