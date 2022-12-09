package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(replaceSpace("We are happy."))
	fmt.Println(replaceSpace("     "))
}

//解法1,直接用库函数
func replaceSpace1(s string) string {
	s = strings.Replace(s, " ", " ", -1)
	return s
}

//解法2，仿造string builder库函数实现
func replaceSpace2(s string) string {
	fmt.Println(s[0:2])

	old := " "
	count := 0
	//计算需要替换的空间
	n := strings.Count(s, old)

	var b strings.Builder //定义一个string builder
	b.Grow(len(s) + count) //预设容量
	start := 0
	for i := 0; i < n; i++ {
		j := start
		j += strings.Index(s[start:], old)
		b.WriteString(s[start:j])//将截取的字符串放到builder里
		b.WriteString("%20")//再加入需要替换的字符串
		start = j + len(old)
	}
	b.WriteString(s[start:])

	return b.String()//返回builder累积的字符串

}


//解法3 ，利用中间切片/数组
func replaceSpace(s string) string {
	b := []byte(s) //string转[]byte
	count := 0
	for _,value := range b{
		if value == ' ' {
			count+=2
		}
	}

	replaces:= make([]byte,len(b)+count)
	n :=0
	for _,value := range b{
		if value == ' ' {
			replaces[n] = '%'
			replaces[n+1] = '2'
			replaces[n+2] = '0'
			n = n+2
		} else {
			replaces[n] = value
		}
		n++
	}

	return string(replaces)//[]byte转string
}