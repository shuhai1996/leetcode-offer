package main

import (
	"fmt"
	"leetcode-offer/mylib"
)

func main() {
	node := mylib.ListNode{}
	head := &node
	head.Val = 5
	node1 := mylib.ListNode{}
	node1.Val = 1
	head.Next = &node1

	node2 := mylib.ListNode{}
	node2.Val = 3
	node1.Next = &node2

	fmt.Println(reversePrint(head))

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//简单暴力解法1
func reversePrint1(head *mylib.ListNode) []int {
	p := head
	var arr []int

	i := 0
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
		i++
	}

	ret := make([]int, i)
	for i > 0 {
		ret[i-1] = arr[len(arr)-i]
		i--
	}

	return ret
}

//解法2，递归
func reversePrint2(head *mylib.ListNode) []int {
	if head == nil {
		return []int{}
	}

	return append(reversePrint2(head.Next), head.Val)
}

//解法3，反转链表
func reversePrint(head *mylib.ListNode) []int {
	if head == nil {
		return []int{}
	}
	var ret []int
	head = reverseLink(head)
	for head != nil{
		ret = append(ret, head.Val)
		head = head.Next
	}

	return ret
}

//利用虚指针双指针反转链表
func reverseLink1(head *mylib.ListNode) *mylib.ListNode {
	var pre *mylib.ListNode
	current := head
	for current != nil {
		temp := current.Next//临时保存temp结点
		current.Next = pre
		pre = current
		current = temp
	}

	return pre
}

//递归反转链表
func reverseLink(head *mylib.ListNode) *mylib.ListNode {
	var pre *mylib.ListNode
	current := head

	return reverse(current, pre)
}

func reverse(cur *mylib.ListNode, pre *mylib.ListNode) *mylib.ListNode{
	if cur == nil{
		return pre
	}

	temp := cur.Next//临时保存temp结点
	cur.Next = pre
	return reverse(temp, cur)
}
