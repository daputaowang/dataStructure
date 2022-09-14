package stack

import (
	"errors"
	"fmt"
)

// 此包定义顺序栈
// 为了方便操作，通常top指示真正的栈顶元素之上的下标地址
// 因为栈底始终在索引为0的位置，所以不需要专门定义
type SqStack struct {
	top       int           // 栈顶索引
	stacksize int           //栈的最大容量
	data      []interface{} // 数据
}

func NewStack(stackSize int) *SqStack {
	// 在初始化结构体时，如果实名初始化字段，不写默认就是该类型的零值，此处为了方便使用按位置初始化，故都写
	return &SqStack{0, stackSize, make([]interface{}, stackSize)}
}

func (s *SqStack) StackEmpty() bool {
	return s.top == 0
}

func (s *SqStack) StackLength() int {
	return s.top
}

func (s *SqStack) ClearStack() {
	s.top = 0
}

func (s *SqStack) DestoryStack() {
	s = nil
}

// 入栈
func (s *SqStack) Push(e interface{}) error {
	if s.top == s.stacksize {
		return errors.New("stack overflow")
	}
	s.data[s.top] = e
	s.top++
	return nil
}

// 出栈
// 弹出最后一个元素并删除
func (s *SqStack) Pop() (interface{}, error) {
	if s.top == 0 {
		return nil, errors.New("stack underflow")
	}
	s.top--
	e := s.data[s.top]
	return e, nil
}

// 返回栈顶元素，但不在栈中删除该元素
func (s *SqStack) Peek() (interface{}, error) {
	if s.top == 0 {
		return nil, errors.New("stack underflow")
	}
	return s.data[s.top-1], nil
}

// 打印栈
func (s *SqStack) Print() {
	if s.top == 0 {
		fmt.Println("empty sqstack")
	}

	// 打印到栈顶元素即可
	fmt.Printf("base-> %v ->top\n", s.data[:s.top])
}
