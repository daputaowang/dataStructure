package queue

import (
	"errors"
	"fmt"
)

// 链式队列使用带头结点和尾指针的链表，当用户无法估计所用队列的长度，则宜采用链式队列

// 定义队列结点
type QNode struct {
	data interface{}
	next *QNode // 此对象是递归的
}

// 定义队列
type LinkQueue struct {
	front  *QNode // 队头指针
	rear   *QNode // 队尾指针
	length int    // 队列长度
}

// 链式队列初始化
// func NewLinkQueue() *LinkQueue {
// 	return &LinkQueue{nil, nil, 0}
// }

// 我认为要加头结点的定义
func NewLinkQueue() *LinkQueue {
	q := &QNode{0, nil}
	return &LinkQueue{q, q, 0}
}

// 以下操作都在链表初始化后执行

// 链式队列销毁
func (lq *LinkQueue) DestroyQueue() {
	lq.front = nil
	lq.rear = nil
}

// 链式队列入队
// 入队插入队尾，尾指针向后移动
// 链表不存在溢出的情况
func (lq *LinkQueue) EnQueue(e interface{}) {
	p := &QNode{e, nil}
	lq.rear.next = p
	lq.rear = p
	lq.length++
}

// 链式队列出队
// 出队是在队头，将首元结点出队，若队列为空则返回underflow。首元结点出队后要将原首元结点的指针域置空
// 如果出队的是首元结点，则在出队后要将尾指针指向头指针
// 返回出队元素
func (lq *LinkQueue) DeQueue() (interface{}, error) {
	// 判断队空
	if lq.front == lq.rear {
		return nil, errors.New("linkqueue underflow")
	}

	// 首元结点出队
	p := lq.front.next
	e := p.data
	lq.front.next = p.next

	// 判断出队的是不是首元结点
	if lq.rear == p {
		lq.rear = lq.front
	}

	// 将原首元结点也就是p的指针域置空（防止其他地方可以通过这个结点引用队列）
	p.next = nil

	lq.length--

	return e, nil
}

// 求链式队列的队头元素
// 拿到元素即可
func (lq *LinkQueue) GetHead() (interface{}, error) {
	if lq.front == lq.rear {
		return nil, errors.New("empty linkqueue")
	}

	return lq.front.next.data, nil
}

// 头插法创建链式队列
// 在空链式队列的情况下使用
func (lq *LinkQueue) CreateLinkQueueT(elems ...interface{}) {
	// 依次将数据插入到尾结点的位置
	for i := 0; i < len(elems); i++ {
		p := &QNode{elems[i], nil}
		lq.rear.next = p
		lq.rear = p
	}
}

// 打印队列
func (lq *LinkQueue) PrintLinkQueue() {
	if lq.front == nil && lq.rear == nil {
		fmt.Println("队列已不存在")
		return
	}

	// 从首元结点开始打印，打印到p.next=nil为止
	p := lq.front.next

	fmt.Println("LinkQueue front")
	for i := 1; p != nil; i++ {
		fmt.Printf("#%d p.data: %v    p.next: %p\n", i, p.data, p.next)
		p = p.next
	}
	fmt.Println("LinkQueue rear")
}
