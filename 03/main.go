package main

import (
	"errors"
	"fmt"
	"os"
)

//环形队列结构体
type CircleQueue struct {
	maxSize int    //4
	array   [5]int //数组
	front   int    //指向队列队首
	rear    int    //指向队尾
}

// Push 添加数据
func (q *CircleQueue) Push(val int) (err error) {
	if q.IsFull() {
		return errors.New("queue full")
	}
	//尾指针在队列尾部，但是不包含最后这个元素
	//大概是留下了一个标志位
	q.array[q.rear] = val
	q.rear = (q.rear + 1) % q.maxSize
	return
}

// Pop 取出数据
func (q *CircleQueue) Pop() (val int, err error) {
	if q.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	//pop,front指向队首，并且包含队首元素
	val = q.array[q.front]
	q.front = (q.front + 1) % q.maxSize
	return val, err
}

// IsFull 判断环形队列是否为满
func (q *CircleQueue) IsFull() bool {
	return (q.rear+1)%q.maxSize == q.front
}

// IsEmpty 判断环形队列是否为空
func (q *CircleQueue) IsEmpty() bool {
	return q.front == q.rear
}

// Size 返回队列元素个数
func (q *CircleQueue) Size() int {
	//关键点
	return (q.rear + q.maxSize - q.front) % q.maxSize
}

// ShowQueue 显示队列,找到队首，然后遍历到队尾
func (q *CircleQueue) ShowQueue() {
	fmt.Println("环形队列内容如下：")
	size := q.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	//定义一个辅助变量，指向front
	tmpFront := q.front
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\n", tmpFront, q.array[tmpFront])
		tmpFront = (tmpFront + 1) % q.maxSize
	}
	fmt.Println()
}

func main() {

	circlQueue := &CircleQueue{
		maxSize: 5,
		front:   0,
		rear:    0,
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
			err := circlQueue.Push(val)
			if err != nil {
				fmt.Println("err =", err.Error())
			} else {
				fmt.Println("已经加入到队列")
				fmt.Println()
			}
		case 2:
			val, err := circlQueue.Pop()
			if err != nil {
				fmt.Println("err =", err.Error())
			} else {
				fmt.Printf("从队列中取出了：%d\n", val)
			}
		case 3:
			circlQueue.ShowQueue()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("指令错误，重新输入")
		}
	}
}
