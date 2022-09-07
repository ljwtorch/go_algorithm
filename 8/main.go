package main

import "fmt"

// PersonNode 创建结构体
type PersonNode struct {
	No   int
	Next *PersonNode //指向下一个人的指针
}

// AddNode 创建函数，形成单向环形链表
func AddNode(num int) *PersonNode {
	//num 人的个数	返回环形链表的第一个人的指针
	first := &PersonNode{} //空结点
	temp := &PersonNode{}  //辅助结点
	//判断
	if num < 1 {
		fmt.Println("num的值不符合要求")
		return first
	}
	//循环构建环形链表
	for i := 1; i <= num; i++ {
		person := &PersonNode{No: i}
		//1.第一个人特殊
		if i == 1 {
			first = person //头结点不要动
			temp = person
			temp.Next = first
		} else {
			temp.Next = person
			temp = person
			temp.Next = first //构成环形链表
		}
	}
	return first
}

// ListNode 遍历显示链表中内容
func ListNode(first *PersonNode) (i int) {
	//如果链表为空
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}
	//创建指针帮助遍历
	temp := first
	for {
		fmt.Printf("编号=%d --> ", temp.No)
		//退出条件
		if temp.Next == first {
			i++
			break
		}
		//移动到下一个
		temp = temp.Next
		i++
	}
	return i
}

/*
分析约瑟夫环思路：
1.编写函数，PlayGame
2.使用算法按照要求，留下最后一个人
*/

func PlayGame(first *PersonNode, startNo int, countNum int, i int) {
	//判空|启动序号要小于总数
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}
	if startNo > i {
		fmt.Println("启动序号大于总数了")
		return
	}
	//定义辅助指针，辅助删除结点
	tail := first
	//让tail指向环形链表的最后一个人，非常重要
	for {
		if tail.Next == first { //说明tail到了最后的小孩
			break
		}
		tail = tail.Next
	}
	//first移动到startNo
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	//数countNum下，移除对应的结点
	for {
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("编号=%d结点被删除\n", first.No)
		//进行删除first指向的结点
		first = first.Next
		tail.Next = first
		//如果tail==first说明只剩下一个结点
		if tail == first {
			break
		}
	}
	fmt.Println("最后一个人的编号为 =", first.No)
}

func main() {
	first := AddNode(6)
	i := ListNode(first)
	fmt.Println()
	PlayGame(first, 1, 5, i)
}
