package main

import (
	"fmt"
	"os"
)

// Emp 定义雇员结构体
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// EmpLink 定义雇员链表 不带表头，第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

// HashTable 定义哈希表，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// ShowMe 显示雇员信息
func (e *Emp) ShowMe() {
	fmt.Printf("链表 %d 找到该雇员 %d\n", e.Id%7, e.Id)
}

// Insert 添加员工的方法，保证添加时，编号从小到大
func (e *EmpLink) Insert(emp *Emp) {
	//辅助指针
	cur := e.Head
	//辅助指针 pre 在 cur 前面
	var pre *Emp = nil
	//如果当前的EmpLink就是一个空链表
	if cur == nil {
		e.Head = emp
		return
	}
	//如果不是一个空链表，给emp找到对应的位置并插入
	//思路是cur 和 emp 进行比较，让pre一直在cur的前面
	for {
		if cur != nil {
			if cur.Id > emp.Id {
				//找到位置了
				break
			}
			//保证同步
			pre = cur
			cur = cur.Next
		}
	}
	//退出后，确认是否将emp添加到链表最后
	pre.Next = emp
	cur.Next = emp
}

// FindById 根据id查找对应热雇员，没有返回nil
func (e *EmpLink) FindById(id int) *Emp {
	//辅助指针
	cur := e.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

// Insert 给HashTable插入雇员方法
func (h *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将该雇员添加到哪个链表
	linkNo := h.HashFun(emp.Id)
	//使用对应的链表添加
	h.LinkArr[linkNo].Insert(emp)
}

// HashFun 散列方法
func (h *HashTable) HashFun(id int) int {
	return id % 7
}

// ShowLink 显示当前链表的信息
func (e *EmpLink) ShowLink(no int) {
	if e.Head == nil {
		fmt.Printf("链表 %d 为空\n", no)
		return
	}
	//遍历当前的链表，并显示数据
	cur := e.Head //辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表 %d 雇员id = %d 名字 = %s -->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() //换行处理
}

// ShewAll 显示hashtable所有的内容
func (h *HashTable) ShowAll() {
	for i := 0; i < len(h.LinkArr); i++ {
		h.LinkArr[i].ShowLink(i)
	}
}

// FindById 查找方法
func (h *HashTable) FindById(id int) *Emp {
	//使用散列函数 确定将该雇员添加到哪个链表
	linkNo := h.HashFun(id)
	return h.LinkArr[linkNo].FindById(id)
}

func main() {
	key := ""
	name := ""
	id := 0
	var hashTable HashTable

	for {
		fmt.Println("----------雇员系统菜单----------")
		fmt.Println("1.添加雇员")
		fmt.Println("2.显示雇员")
		fmt.Println("3.查找雇员")
		fmt.Println("4.退出系统")
		fmt.Print("请输入对应编号：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Print("输入雇员id：")
			fmt.Scanln(&id)
			fmt.Print("输入雇员name：")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "2":
			hashTable.ShowAll()
		case "3":
			fmt.Print("输入id号：")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				//显示雇员信息
				emp.ShowMe()
			}
		case "4":
			os.Exit(0)
		default:
			fmt.Println("输入有误，重新输入！")
		}
	}
}
