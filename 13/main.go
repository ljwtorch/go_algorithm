package main

import (
	"errors"
	"fmt"
)

// Stack 使用数组来模拟一个栈的使用
type Stack struct {
	Max int      //栈最大存放个数
	Top int      //表示栈顶
	arr [100]int //数组模拟栈
}

// Push 入栈
func (s *Stack) Push(val int) (err error) {
	//先判断栈是否已满
	if s.Top == s.Max-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	//栈顶自增后加入数据
	s.Top++
	s.arr[s.Top] = val
	return
}

// Pop 出栈
func (s *Stack) Pop() (val int, err error) {
	//判空
	if s.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack full")
	}
	//先取值 栈顶指针再--
	val = s.arr[s.Top]
	s.Top--
	return val, nil
}

// List 栈的遍历
func (s *Stack) List() {
	//判空
	if s.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈中数据为：")
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, s.arr[i])
	}
}

func main() {
	stack := Stack{
		Max: 100,
		Top: -1,
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push(8)

	stack.List()
	fmt.Println("数据出栈：")
	for i := 0; i < 8; i++ {
		val, _ := stack.Pop()
		if val != 0 {
			fmt.Printf("Pop=%d\n", val)
		}
	}
	stack.List()
}
