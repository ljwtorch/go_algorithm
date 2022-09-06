package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
1、数组模拟循环队列，每隔一段时间[随机的]，给数组添加一个数。
2、启动两个协程，每隔一定时间，到队列取出数据。
3、在控制台输出
4、使用锁机制
*/

// CircleQueue 创建一个数组结构体
type CircleQueue struct {
	maxSize int
	array   [50]int
	front   int
	rear    int
}

//定义一个全局互斥锁
var lock sync.Mutex

// Push 插入数据
func (c *CircleQueue) Push(n int) (err error) {
	//判断队列是否为满
	if c.IsFull() {
		return errors.New("queue full")
	}
	//队尾插入数据
	c.array[c.rear] = n
	c.rear = (c.rear + 1) % c.maxSize
	return
}

// Pop 取出数据并查看结果
func (c *CircleQueue) Pop(n int) {
	for {
		if c.IsEmpty() {
			//设置延迟取出数据
			time.Sleep(time.Second)
			fmt.Println("没有人在排队……")
		} else {
			//Duration表示两个时间点之间经过的时间/Intn返回区间内的伪随机数
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			//加锁 避免竟态
			lock.Lock()
			//赋值元素n，front指向队首，包含队首元素
			val := c.array[c.front]
			//打印结果
			fmt.Printf("%d号协程服务 ==> %d号客户\n", n, val)
			//取出当前元素，头结点后移
			c.front = (c.front + 1) % c.maxSize
			lock.Unlock()
		}
	}
}

// IsFull 判满
func (c *CircleQueue) IsFull() bool {
	return (c.rear+1)%c.maxSize == c.front
}

// IsEmpty 判空
func (c *CircleQueue) IsEmpty() bool {
	return c.front == c.rear
}

func main() {
	//初始化结构体
	circleQueue := &CircleQueue{
		maxSize: 50,
		front:   0,
		rear:    0,
	}
	//创建随机
	rand.Seed(time.Now().Unix())
	temp := -1
	//启动两个协程，循环读取数据
	go circleQueue.Pop(1)
	go circleQueue.Pop(2)
	//循环放入数据
	for {
		temp = rand.Intn(100)
		//fmt.Println(temp)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		circleQueue.Push(temp)
		circleQueue.Push(temp)
	}
}
