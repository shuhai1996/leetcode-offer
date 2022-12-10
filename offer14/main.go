package main

import "fmt"

func main()  {
	//经典的动态规划题
	fmt.Println(cuttingRope(8))
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
		dp := make([]int, n+1)//结果放入切片，0，1，2，3只是为了计算
		dp[0] = 0
		dp[1] = 1
		dp[2] = 2
		dp[3] = 3
		for i := 4; i <= n; i++ {
			max := 0
			for j := 1; j <= i/2; j++ {
				temp := dp[j]*dp[i-j]// f(n) = max(f(i)*f(n-i))
				if max < temp {
					max = temp
				}
				dp[i] = max
			}
		}

		return dp[n]
	}
}