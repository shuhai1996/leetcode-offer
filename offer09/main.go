package main

import "fmt"

func main()  {

	s1 := []int{1,2,3,4}
	var s2 []int
	for _, v := range s1 {
		s1 = s1[:len(s1)-1]
		fmt.Println(v)
	}

	fmt.Println(s1)
	fmt.Println(s2)


	obj := Constructor()
	obj.AppendTail(1)
	obj.AppendTail(3)
	param1 := obj.DeleteHead()
	param2 := obj.DeleteHead()
	obj.AppendTail(5)
	obj.AppendTail(2)
	param3 := obj.DeleteHead()

	param4 := obj.DeleteHead()
	fmt.Println(param1, param2,param3,param4)
}

type CQueue struct {
	inStack,outStack []int //定义进栈 和出栈
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
	this.inStack = append(this.inStack, value)//插入队尾元素写入进栈中
}


func (this *CQueue) DeleteHead() int {
	var value int
	//fmt.Println(this.outStack)
	if len(this.outStack) == 0 {//如果出栈里没有元素
		if len(this.inStack) == 0 {//进栈里也没有元素返回-1
			return -1
		} else { //进栈当中与有元素，则依次写入到出栈中，每次删除进栈的前一个元素
			for _, v := range this.inStack { //遍历进栈，循环过程中改变切片不会对遍历数据造成影响
				this.inStack = this.inStack[:len(this.inStack)-1]//切片长度-1
				this.outStack = append(this.outStack, v)//放入到出栈中
			}
			//fmt.Println(this.inStack)
			return this.DeleteHead()
		}
	} else {
		value = this.outStack[0]//返回出栈中第一个元素
		if len(this.outStack) == 1 {
			this.outStack = []int{}
		} else {
			this.outStack = this.outStack[1:]
		}
		return value
	}

}


