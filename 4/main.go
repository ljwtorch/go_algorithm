package main

import "fmt"

type HeroNode struct {
	No       int
	Name     string
	NickName string
	next     *HeroNode //指向下一个结点
}

// Insert 给链表从插入结点(无序插入)
func Insert(head *HeroNode, newNode *HeroNode) {
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

// InsertNode 有序插入方法，根据No编号从小到大插入
func InsertNode(head *HeroNode, newNode *HeroNode) {
	//1.找到适当的结点
	//2.创建一个辅助结点
	temp := head
	flag := true
	//让插入结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.No > newNode.No { //改符号方向可变升序或降序 改成等于新结点会排在同结点的前面
			//说明新节点应该插入到temp后面
			break
		} else if temp.next.No == newNode.No {
			//说明链表中已经有这个No，不插入
			flag = false //标识符变成false 不可插入
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("已经存在的No =", newNode.No)
		return
	} else {
		//先让新结点连，防止断链
		newNode.next = temp.next
		temp.next = newNode
	}
}

//删除结点
func DeleteNode(head *HeroNode, id int) {
	//1.创建一个辅助结点
	temp := head
	flag := false
	//找到要删除的结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.No == id {
			//说明找到了
			flag = true
			break
		}
		temp = temp.next
	}
	//找到后删除
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("要删除的结点不存在")
	}
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

	hero3 := &HeroNode{
		No:       3,
		Name:     "沙悟净",
		NickName: "沙僧",
	}

	//3.加入
	InsertNode(head, hero3)
	InsertNode(head, hero1)
	InsertNode(head, hero2)

	//4.显示
	fmt.Println("插入后的链表：")
	ListNode(head)

	//5.删除
	fmt.Println()
	fmt.Println("删除后的链表：")
	DeleteNode(head, 2)
	ListNode(head)
}
