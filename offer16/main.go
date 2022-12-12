package main

import (
	"fmt"
)

func main() {

	//fmt.Println(20>>1)

	//fmt.Println(math.Pow(1, -2147483648))
	fmt.Println(myPow(2.0 , -2))
	fmt.Println(myPow(0.001 , 2147483647))
}

func myPow(x float64, n int) float64 {
	if x == 1 {
		return 1

	}

	if n >= 0 {
		return cal(x, n)
	}

	if n <0 {
		return 1.0/cal(x, -n)
	}

	return 0
}

func cal(x float64, n int)  float64{
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	ret := myPow(x, n>>1)//右移运算代替/2
	ret *= ret
	if n % 2 == 1 {
		ret *= x
	}

	return ret
}

