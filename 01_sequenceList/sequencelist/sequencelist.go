package sequencelist

import (
	"errors"
	"fmt"
)

// 1、顺序表在golang中使用切片作为底层数据结构
// 2、顺序表的取值、插入、删除都是按照slice下标进行访问的。也就是随机访问。
// 3、数组元素是int类型
// 4、本例中的按位置删除传入的index直接与下标对应，不是序号
// 5、Author: daputaowang

// 定义一个顺序表
type SqList struct {
	elem   []int
	length uint
}

// 为数组初始化内存
func NewSqList(capacity uint) *SqList {
	if capacity == 0 {
		return nil
	}

	return &SqList{
		elem:   make([]int, capacity),
		length: 0,
	}
}

// 销毁:golang中不需要手动释放空间，会存在GC

// 获取长度
func (s *SqList) len() uint {
	return s.length
}

// 判空
func (s *SqList) isEmpty() bool {
	return s.length == 0
}

// 判断索引是否越界
func (s *SqList) isIndexOutOfRange(index uint) bool {
	return index >= uint(cap(s.elem))
}

// 按下标取值
func (s *SqList) GetElem(index uint) (int, error) {
	if s.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return s.elem[index], nil
}

// 查找，找到该数据的位置（从0开始的位置）
func (s *SqList) LocateElem(e int) (int, error) {
	// 引入哨兵思想
	if s.isEmpty() {
		return 0, errors.New("SqList is empty")
	}

	if temp := s.elem[s.length-1]; temp == e {
		return int(s.length), nil
	}

	s.elem[s.length-1] = e

	locate := 0
	for i := 0; ; i++ {
		if s.elem[i] == e {
			locate = i + 1
			break
		}
	}

	if locate == int(s.length) {
		return 0, errors.New("the element is not in the SqList")
	}

	return locate, nil
}

// 按下标插入
func (s *SqList) Insert(index uint, v int) error {
	if s.len() == uint(cap(s.elem)) {
		return errors.New("full SqList")
	}

	if s.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	// 在index位置插入数据，插入后后面的元素全部需要后移
	// 因为是后移所以从最后一个位置移动方便
	for i := s.length - 1; i >= index; i-- {
		s.elem[i+1] = s.elem[i]
	}
	s.elem[index] = v
	s.length++
	return nil
}

// 按位置删除，并返回删除的值
func (s *SqList) Delete(index uint) (elem int, err error) {
	if s.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}

	// 将index位置的数据先放入返回值中
	// 将index之后的数据都向前移动一位
	// 循环条件只需要到length-2次也就是i<s.length-1，<就会使最后一个下标不符合条件，因为在最后一次移动中，最后一个elem已经被移动了
	elem = s.elem[index]
	for i := index; i < s.length-1; i++ {
		s.elem[i] = s.elem[i+1]
	}
	s.length--
	return elem, nil
}

func (s *SqList) Print() {
	var fmtStr string
	for i := uint(0); i < s.len(); i++ {
		fmtStr += fmt.Sprintf("%v ", s.elem[i])
	}
	fmt.Println(fmtStr)
}
