package main

import "fmt"

func main() {
	//抓住非递减条件，逐步缩减查找范围
	//这个题需要注意边界问题！
	//fmt.Println(findNumberIn2DArray([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26,30}}, 27))
	//fmt.Println(findNumberIn2DArray([][]int{}, 0))
	fmt.Println(findNumberIn2DArray([][]int{{-5}, {-10}}, -10))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	ret := false
	rows := len(matrix) //二维数组行数
	if rows == 0 {
		return false
	}
	cols := len(matrix[0]) //二维数组列数。注意边界
	if cols == 0 {
		return false
	}

	if matrix[rows-1][cols-1] >= target { //同样需要注意边界
		r := 0
		c := cols - 1
		for c >= 0 && r <= rows-1 { //循环终止的条件
			if matrix[r][c] == target {
				ret = true
				break
			} else if matrix[r][c] > target { //大于target消去一列
				c--
			} else { //小于target往下移动一行
				r++
			}
		}
	}

	return ret
}
