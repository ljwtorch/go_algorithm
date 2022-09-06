package main

import "fmt"

type HeroNode struct {
	No       int
	Name     string
	NickName string
	next     *HeroNode //指向下一个结点
}

// InsertNode 给链表从插入结点 编写第一种插入方法，在单链表的最后加入
func InsertNode(head *HeroNode, newNode *HeroNode) {
	//1.先找到链表的最后结点
	//2.创建一个辅助结点
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next //让temp不断指向下一个结点
	}
	//3。将newNode加入到链表的最后
	temp.next = newNode
}

// ListNode 显示链表的所有结点信息
func ListNode(head *HeroNode) {
	//1.创建一个辅助结点
	temp := head

	//2.判断链表是不是空链表
	if temp.next == nil {
		fmt.Println("链表为空！")
		return
	}

	//3.遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.No, temp.next.Name, temp.next.NickName)
		//判断是否在链表最后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	//1.创建头结点
	head := &HeroNode{}

	//2.创建新的heroNode
	hero1 := &HeroNode{
		No:       1,
		Name:     "宋江",
		NickName: "及时雨",
	}

	hero2 := &HeroNode{
		No:       2,
		Name:     "孙悟空",
		NickName: "美猴王",
	}

	//3.加入
	InsertNode(head, hero1)
	InsertNode(head, hero2)
	//4.显示
	ListNode(head)
}
