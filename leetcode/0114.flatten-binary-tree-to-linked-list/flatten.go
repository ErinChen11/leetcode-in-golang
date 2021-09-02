package flattenbinarytreetolinkedlist

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}

	// 對左右子樹調用flatten
	flatten(root.Left)
	flatten(root.Right)

	// 先將左右子樹先拉平， 複製好左右子樹
	left := root.Left 
	right := root.Right
	// 將左子樹轉為右子樹, 原本的左子樹為空
	root.Left = nil
	root.Right = left
	// 再將右子樹接到當前右子樹下面
	for root.Right != nil {
		// 如果新的右子樹有東西, 就一直往下接
		root = root.Right
	}
	// root 一直往下接, 接到 root.Right 是nil 沒有東西了
    // 表示新的右子樹已經到盡頭, 就把舊的右子樹right接上去
	root.Right = right




}

