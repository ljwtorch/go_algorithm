package main

import (
	"errors"
	"fmt"
	"os"
)

/*
存入队列有两个步骤：
1.将尾指针向后移：rear+1,front == rear(空)
2.若尾指针rear小于等于队列的最大下标MaxSize-1,将数据存入rear所指的数组元素中，否则无法存入数据
*/

type Queue struct {
	maxSize int
	array   [5]int //模拟队列数组
	front   int    //表示指向队列
	rear    int    //表示指向队列的尾部
}

// AddQueue 添加数据到队列
func (q *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if q.rear == q.maxSize-1 {
		return errors.New("queue full")
	}
	q.rear++ //rear后移
	q.array[q.rear] = val
	return
}

// ShowQueue 显示队列,找到队首，然后遍历到队尾
func (q *Queue) ShowQueue() {
	fmt.Println("当前队列的情况是：")
	//q.front是不包含队首的元素
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d  ", i, q.array[i])
	}
	fmt.Println()
	fmt.Println()
}

// GetQueue 取出数据
func (q *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if q.front == q.rear {
		return -1, errors.New("queue empty")
	}
	q.front++
	//fmt.Println(q.front)
	val = q.array[q.front]
	return val, err
}

func main() {
	/*
		使用数组实现队列的思路
		1.创建一个数组array，是作为队列的一个字段
		2.front初始化为-1
		3.rear表示队列尾部，初始化为-1
		4.完成队列的基本查找	AddQueue	GetQueue	ShowQueue
	*/
	//先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key int
	var val int
	for {
		fmt.Println("1.添加数据")
		fmt.Println("2.获取数据")
		fmt.Println("3.显示数据")
		fmt.Println("0.退出")
		fmt.Print("输入1-4：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Print("输入要入队列的数:")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println("err =", err.Error())
			} else {
				fmt.Println("已经加入到队列")
				fmt.Println()
			}
		case 2:
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println("err =", err.Error())
			} else {
				fmt.Printf("从队列中取出了：%d\n", val)
			}
		case 3:
			queue.ShowQueue()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("指令错误，重新输入")
		}
	}

}
