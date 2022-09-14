package stack

import (
	"errors"
	"fmt"
)

// 此包定义链栈
// 链栈没有栈满的情况
// 链栈使用没有头结点的单链表，栈顶在头指针指向的地方，栈底就是链表的尾部，栈底的结点的指针域为nil
// 链栈入栈是向头指针指向的位置也就是栈顶插入结点，出栈也是对头指针指向的结点进行操作
// 空栈相当于头指针指向空
type StackNode struct {
	data interface{}
	next *StackNode
}

// 如果不构造这一步，直接使用new(StackNode)，在方法中使用时想直接修改头指针会很困难
// 因为方法中传递指针类型的接收器作为形参对指针的拷贝，直接修改指针不能产生引用传递的效果
type LinkStack struct {
	top *StackNode
}

func NewLinkStack() *LinkStack {
	// 构造一个空栈，栈顶指针置为空
	return &LinkStack{new(StackNode)}
}

func (s *LinkStack) StackEmpty() bool {
	return s.top.data == nil
}

// 入栈
func (s *LinkStack) Push(e interface{}) {
	// 生成新节点p，数据域置为e，指针域指向s也就是栈顶
	if s.top.data == nil {
		// 表示插入的结点是链表第一个结点
		s.top.data = e
		return
	}
	p := &StackNode{e, s.top}
	s.top = p
}

// 出栈
func (s *LinkStack) Pop() (interface{}, error) {
	if s.top.data == nil {
		return nil, errors.New("linkstack underflow")
	}
	e := s.top.data
	p := s.top
	s.top = s.top.next
	// 把原栈顶结点的next域置为nil相当于删除原栈顶结点
	p.next = nil
	return e, nil
}

// 返回栈顶元素，不删除该结点
func (s *LinkStack) Peek() interface{} {
	if s.top.data != nil {
		return s.top.data
	} else {
		return nil
	}
}

// 尾插法批量创建链栈
func (s *LinkStack) CreateLinkT(elems ...interface{}) {
	// 依次将数据插入到尾结点的位置
	tail := s.top
	for i := len(elems); i > 0; i-- {
		if tail.data == nil {
			tail.data = elems[i-1]
			continue
		}
		p := &StackNode{elems[i-1], nil}
		tail.next = p
		tail = p
	}
}

// 打印链栈
func (s *LinkStack) PrintLink() {
	// 从首元结点开始打印，打印到p.next=nil为止
	p := s.top

	fmt.Println("stack front")
	for i := 1; p != nil; i++ {
		fmt.Printf("#%d p.data: %v    p.next: %p\n", i, p.data, p.next)
		p = p.next
	}
	fmt.Println("stack rear")
}
