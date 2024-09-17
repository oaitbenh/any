package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	arr := []int{4, 11, 2}
	FirstNode := TreeNode{
		Val: 5,
	}
	Cur := FirstNode
	for i := 0; i < len(arr); i++ {
		Cur.Right.Val = arr[i]
		if i%2 == 0 {
			Cur.Right = &TreeNode{}
			Cur = *Cur.Right
		} else {
			Cur.Left = &TreeNode{}
			Cur = *Cur.Left
		}
	}
}
