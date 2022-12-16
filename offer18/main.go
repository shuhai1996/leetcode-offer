package main

import (
	"fmt"
	"leetcode-offer/mylib"
)

func main()  {
	head := mylib.ListNode{Val: 1}
	node1 := mylib.ListNode{Val: 3}
	head.Next = &node1
	node2 := mylib.ListNode{Val: 3}
	node1.Next = &node2
	node3 := mylib.ListNode{Val: 4}
	node2.Next = &node3

	p := deleteNode(&head, 3)
	fmt.Println(p)
	for p!=nil{
		fmt.Println(p.Val)
		p = p.Next
	}
}

func deleteNode(head *mylib.ListNode, val int) *mylib.ListNode {
	temp := &mylib.ListNode{Val: 0, Next: head}//临时指针，next指向头部
	p := temp
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		}
		p = p.Next
	}

	return temp.Next
}