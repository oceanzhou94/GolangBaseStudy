package main

import (
	"errors"
	"fmt"
)

// Queue 定义队列结构体
type Queue struct {
	data  []int // 队列的数据存储
	front int   // 队头指针
	rear  int   // 队尾指针
}

// Init 初始化队列
func (q *Queue) Init() {
	q.data = make([]int, 0)
	q.front = -1
	q.rear = -1
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.front == -1
}

// Length 获取队列长度
func (q *Queue) Length() int {
	if q.IsEmpty() {
		return 0
	}
	return q.rear - q.front + 1
}

// Front 获取队头元素
func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	return q.data[q.front], nil
}

// Rear 获取队尾元素
func (q *Queue) Rear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	return q.data[q.rear], nil
}

// Enqueue 入队
func (q *Queue) Enqueue(value int) {
	if q.IsEmpty() {
		q.front++
		q.rear++
		q.data = append(q.data, value)
	} else {
		q.rear++
		q.data = append(q.data, value)
	}
}

// Dequeue 出队
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	value := q.data[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front++
	}
	return value, nil
}

// Print 打印队列
func (q *Queue) Print() {
	for i := q.front; i <= q.rear; i++ {
		fmt.Printf("%d ", q.data[i])
	}
	fmt.Println()
}

func main() {
	// 创建一个队列对象
	q := Queue{}
	q.Init()

	// 入队
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// 打印队列
	q.Print()

	// 出队
	value, _ := q.Dequeue()
	fmt.Printf("Dequeue %d\n", value)

	// 打印队列
	q.Print()

	// 获取队头元素
	value, _ = q.Front()
	fmt.Printf("Front %d\n", value)

	// 获取队尾元素
	value, _ = q.Rear()
	fmt.Printf("Rear %d\n", value)

	// 获取队列长度
	fmt.Println("Queue length is", q.Length())
}
