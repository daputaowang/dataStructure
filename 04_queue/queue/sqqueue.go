package queue

import (
	"errors"
	"fmt"
)

// 顺序队列
// 循环队列解决假溢出
// 队空与队满的解决方案采用空出一个元素空间法：尾指针指向队列尾部+1的位置，用于区分队满（(rear+1)%capacity==front）与队空（rear == front）
type SqQueue struct {
	s        []interface{} // 实际存储队列的slice
	capacity int           // 队列的容量
	front    int           // 头指针
	rear     int           // 尾指针，尾部+1的位置
}

// 队列的初始化
func NewSqQueue(capacity int) *SqQueue {
	// 开始的时候0位置没有值，可以想象队尾的最后一个元素在-1的位置，所以符合尾指针在尾部+1的规则
	return &SqQueue{make([]interface{}, capacity), capacity, 0, 0}
}

// 求队列的长度
// 因为是循环队列，如果进入循环之后队尾指针因为取余的原因会小于队头的元素，而此时的队尾指针实际上相当于是加上一个容量的长度
// 为了计算方便，未进入循环的情况也加上一个容量的长度，通过取余一个容量的长度，就可以将多加的化解掉
func (q *SqQueue) QueueLength() int {
	return (q.rear + q.capacity - q.front) % q.capacity
}

// 循环队列入队
// 入队是在队尾入队，因为rear指向队尾加一的地方，所以先判断队满，没满的话将新元素加入当前队尾位置，然后队尾指针循环加一也就是加一取余
func (q *SqQueue) EnQueue(e interface{}) error {
	// 判断队满
	if (q.rear+1)%q.capacity == q.front {
		return errors.New("sqqueue overflow")
	}

	// 插入新元素
	q.s[q.rear] = e

	// 尾指针循环加一
	q.rear = (q.rear + 1) % q.capacity

	return nil
}

// 循环队列出队
// 出队是在队头的位置出队，头指针直接指向队头，先判断队空，没空的话将队头元素出队，队头指针循环加一
// 返回要出队的元素
func (q *SqQueue) DeQueue() (interface{}, error) {
	// 判断队空
	if q.front == q.rear {
		return nil, errors.New("sqqueue underflow")
	}

	// 取出出队元素
	e := q.s[q.front]

	// 头指针循环加一
	q.front = (q.front + 1) % q.capacity

	return e, nil
}

// 循环队列取队头元素
func (q *SqQueue) GetHead() (interface{}, error) {
	// 判断队空
	if q.front == q.rear {
		return nil, errors.New("sqqueue underflow")
	}

	// 取出出队元素
	e := q.s[q.front]
	return e, nil
}

// 打印队列
func (q *SqQueue) Print() {
	if q.front == 0 && q.rear == 0 {
		fmt.Println(nil)
	} else if q.front < q.rear {
		fmt.Printf("front-->  %v  -->rear", q.s[q.front:q.rear]) //左开右闭
	} else if q.front > q.rear {
		// 插入的数据进行了取余循环
		temp := append(q.s[q.front:], q.s[:q.rear]...)
		fmt.Printf("front-->  %v  -->rear", temp)
	}
}
