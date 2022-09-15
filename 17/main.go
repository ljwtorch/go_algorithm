package main

import "fmt"

type Tree struct {
	No    int
	Name  string
	Left  *Tree
	Right *Tree
}

// PreOrder 前序遍历：根左右
func PreOrder(node *Tree) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// InfixOrder 中序遍历：左根右
func InfixOrder(node *Tree) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

// PostOrder 后序遍历：左右根
func PostOrder(node *Tree) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func main() {
	//构建二叉树
	root := &Tree{
		No:   1,
		Name: "one",
	}

	left1 := &Tree{
		No:   2,
		Name: "two",
	}

	right1 := &Tree{
		No:   3,
		Name: "three",
	}

	right2 := &Tree{
		No:   4,
		Name: "four",
	}

	root.Left = left1
	root.Right = right1
	right1.Right = right2

	fmt.Println("-----前序遍历-----")
	PreOrder(root)
	fmt.Println("-----中序遍历-----")
	InfixOrder(root)
	fmt.Println("-----后序遍历-----")
	PostOrder(root)
}
