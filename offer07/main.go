package main

import (
	"fmt"
	"leetcode-offer/mylib"
)

func main()  {
	fmt.Println(buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7}))
	fmt.Println(mylib.InOrderTree(buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})))
}

func buildTree(preorder []int, inorder []int) *mylib.TreeNode{
	if len(preorder) ==0 || len(inorder) == 0 {
		return nil
	}

	root := &mylib.TreeNode{
		Val: preorder[0],
	}

	//先通过中序遍历找出根结点位置，即可分离出左右子树
	i := 0
	for index, value:=range inorder{
		if value == root.Val {
			i = index
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i]) //  inorder[:i] 就是左子树中序遍历的序列, 那么可以推算出其前序遍历序列为preorder[1:len(inorder[:i])+1]，同理可以推算出右子树结构
	root.Right = buildTree(preorder[len(inorder[:i])+1:],inorder[i+1:])
	return root
}

