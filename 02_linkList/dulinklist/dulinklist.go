package dulinklist

import "errors"

// 初始化双向链表
type Elem int

type DuLinkNode struct {
	data  Elem
	prior *DuLinkNode
	next  *DuLinkNode
}

// 双向链表的插入
// 1.先找到要插入的位置i∈[1,n+1]，因为是双向链表所以直接找到第i个结点即可，不需要找i-1
// 2.i∈[1,n+1]，可以插入到表长+1的地方，所以空表的第1个位置可插入，表长+1的位置可以插入，超过表长的位置不能插入。
// 3.构造要插入的结点S，S的数据域是d，S的前驱结点是第i个结点的前驱结点，后置结点就是第i个结点
func (head *DuLinkNode) Insert(i int, d Elem) error {
	p := head
	j := 0
	for p.next != nil && j < i {
		p = p.next
		j++ // 写在此处便于理解
	}

	if (p.next == nil && j < i-1) || j > i-1 { // p == nil处理i大于length+1的情况，j > i-1处理i等于0和小于0的情况，所以i可以是int
		return errors.New("i is not in the [1, length+1] range")
	}

	// 初始化s
	s := &DuLinkNode{d, nil, nil}
	// 处理给空表插入和在表长+1插入的情况
	if p.next == nil && j == i-1 {
		// s的前驱结点为p，p的后继结点为s
		s.prior = p
		p.next = s
		// s的后继结点为nil
		s.next = nil
	}

	// p的前一个结点作为s的前驱，s作为前一个结点的后继
	s.prior = p.prior
	p.prior.next = s
	// p作为s的后继，s作为p的前驱
	s.next = p
	p.prior = s
	return nil
}

// 双向链表的删除
// 1.先找到要删除的位置i∈[1,n]，因为是双向链表所以直接找到第i个结点即可，不需要找i-1
// 2.i可能小于1，可能大于表长，这都是非法的
// 3.返回删除的结点
func (head *DuLinkNode) Delete(i int) (*DuLinkNode, error) {
	p := head.next
	j := 1
	for ; p != nil && j < i; j++ {
		p = p.next
	}

	if p == nil || j > i {
		return nil, errors.New("i is not in the [1, length] range")
	}
	q := p
	// p的前驱结点的后继变为p的后继结点
	p.prior.next = p.next
	// 如果p的后继结点不为nil，p的后继结点的前驱结点变为p的前驱结点
	if p.next != nil {
		p.next.prior = p.prior
	}
	return q, nil
}
