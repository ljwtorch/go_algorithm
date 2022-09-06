package main

import "fmt"

// CircleNode 定义结构体结点
type CircleNode struct {
	No   int
	Name string
	next *CircleNode
}

// InsertNode 插入结点
func InsertNode(head *CircleNode, newNode *CircleNode) {
	//1.判断是不是添加第一个结点
	if head.next == nil {
		head.No = newNode.No
		head.Name = newNode.Name
		head.next = head //构成环形
		fmt.Println("结点已经插入！")
		return
	}

	//2.定义临时变量，找到最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = newNode
	newNode.next = head
}

// ListNode 显示链表内容
func ListNode(head *CircleNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空！")
		return
	}
	for {
		fmt.Printf("[%d , %s]==>", temp.No, temp.Name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}

// DeleteNode 删除结点
func DeleteNode(head *CircleNode, id int) *CircleNode {
	/*
		删除思路：
		1.先让temp指向head
		2.让helper指向环形链表最后
		3.让temp和要删除的id进行比较，如果相同则通过helper完成删除（还要考虑如果删除的就是头结点怎么办）
	*/
	temp := head
	helper := head

	//空链表的情况
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，不能删除！")
		return head
	}

	//如果只有一个结点
	if temp.next == head {
		if temp.No == id {
			temp.next = nil
		}
		return head
	}

	//将helper定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//如果有两个及以上的结点
	flag := true
	for {
		if temp.next == head { //如果为true，说明已经比较到最后一个，最后一个还没有比较
			break
		}
		if temp.No == id { //找到对应id
			if temp == head { //说明删除的是头结点
				head = head.next
			}
			helper.next = temp.next
			fmt.Println("删除了结点:", id)
			flag = false
			break
		}
		temp = temp.next     //这个移动是为了比较
		helper = helper.next //这个移动是为了删除
	}
	//还剩一次没有比较
	if flag { //如果flag为true，说明上边没有删除结点
		if temp.No == id {
			helper.next = temp.next
			fmt.Println("删除了结点:", id)
		} else {
			fmt.Println("不存在结点:", id)
		}
	}
	return head
}

func main() {
	//初始化一个环形链表的头结点
	head := &CircleNode{}

	//创建结点
	node1 := &CircleNode{
		No:   1,
		Name: "one",
	}
	node2 := &CircleNode{
		No:   2,
		Name: "two",
	}
	node3 := &CircleNode{
		No:   3,
		Name: "three",
	}
	node4 := &CircleNode{
		No:   4,
		Name: "four",
	}

	//插入结点
	InsertNode(head, node1)
	InsertNode(head, node2)
	InsertNode(head, node3)
	InsertNode(head, node4)

	ListNode(head)

	DeleteNode(head, 2)

	ListNode(head)

}
