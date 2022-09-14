package linklist

import (
	"errors"
	"fmt"
)

// 此包中为带头结点的链表
// 初始化的第一个结点为头结点，头结点的指针即为头指针，头结点的数据域不存储内容

// 单链表部分

// 定义结点

type Elem int
type LinkNode struct {
	data Elem
	next *LinkNode
}

// 初始化单链表，头结点的数据域存什么都行，指针域存储nil。返回的头结点的指针就代指链表
func NewLinkNode() *LinkNode {
	return &LinkNode{0, nil}
}

// 链表的几个重要基本操作
// p = head 使p指向头结点
// q = p.head 头结点的next就是首元结点的地址，也就指向了首元结点
// p = p.next p指向下一个结点

// 单链表判空
func (head *LinkNode) ListEmpty() bool {
	// 判断头结点中的指针域的值是否为nil
	return head.next == nil
}

// 单链表销毁
// O(n)
func (head *LinkNode) DestroyLinkList() {
	// 此处思考一个问题golang中是否可以通过直接将头指针置位nil，使后续链表没有得到引用，从而触发gc自动回收（）
	// 销毁的思想是先让头指针指向下一个结点，然后删除当前结点
	// 结束条件是l == nil 循环条件是l != nil

	var p *LinkNode
	for head != nil {
		// 先将当前头结点的下一个结点保存地址，以便于删除当前头结点后，后续结点的寻找
		p = head.next
		// 然后将当结点置空
		head = nil
		// 然后将head向后移动
		head = p
	}

}

// 单链表清空
// O(n)
func (head *LinkNode) ClearLinkList() {
	// 此处思考一个问题，golang清空链表是否可以直接将头结点的指针域置为nil，后续链表部分由GC回收
	// 清空的思想是只保留头指针和头结点，链表中的元素全部删除，
	// 相比于销毁链表而言就是将从头结点开始删除，变为从首元结点开始删除，也就是将head变为head.next
	// 先用临时变量记录首元结点的指针也就是head.next，然后再使用一个临时变量保存
	var p, q *LinkNode
	p = head.next
	// for {
	// 	if p == nil {
	// 		break
	// 	}
	// 	q = p.next
	// 	p.next = nil
	// 	p = q
	// }

	for p != nil {
		q = p.next
		p.next = nil
		p = q
	}

	// 将头结点的指针域置为nil
	head.next = nil
}

// 单链表表长
// O(n)
func (head *LinkNode) LinkListLen() uint {
	// 求单链表的表长为从首元结点开始计数，直到某个结点的next == nil为止

	// 按照c中的习惯，一个结点这样定义
	// 但在golang中如果有类型确定值，直接可以使用p := head.next 来代替下面两行，此处第一次出现暂时这样写，后面全部简写
	var p *LinkNode
	p = head.next

	if p == nil {
		return 0
	}
	count := uint(0)
	for p != nil {
		count++
		p = p.next
	}
	return count
}

// 单链表取值：取单链表中第i∈[1,n]个元素的内容，使用uint避免出现负值
// 步骤：从链表的头指针出发，顺着next一直往下搜索，直到第i个结点为止。取出data
// O(n)
func (head *LinkNode) GetElem(i int) (Elem, error) {
	p := head.next
	// 链表中的位置相关的循环从1开始，便于思考
	// 因为每次循环p都会指向下一个结点，所以这种在循环体中指向下一个结点的情况循环n-1次就可以了，也就是j<i
	// 当p为nil时说明上一次循环中获取的p.next取到的是nil，则链表已经遍历结束，传入的位置超过元素的长度。所以不取

	j := 1
	for ; p != nil && j < i; j++ {
		p = p.next
	}

	if p == nil || j > i {
		return 0, errors.New("out of length range")
	}

	return p.data, nil
}

// 单链表按值查找：根据指定数据获取数据所在的位置(位置指针)
// 链表查找的循环条件为 p!=nil p.data != data，前者判断链表遍历完毕元素没有找到的情况
func (head *LinkNode) LocateElem(data Elem) (*LinkNode, error) {
	// 数据有可能不在链表中，不在则返回error。不在链表中表现为p == nil时元素依然没有找到
	p := head.next // 指向首元结点
	for p != nil && p.data != data {
		p = p.next
	}
	if p == nil {
		return nil, errors.New("the element is not in the linklist")
	}
	return p, nil
}

// 单链表按值查找：根据指定数据获取数据所在的位置序号
func (head *LinkNode) LocateElemIDX(data Elem) (uint, error) {
	// 数据有可能不在链表中，不在则返回error。不在链表中表现为p == nil时元素依然没有找到
	p := head.next // 指向首元结点
	i := uint(1)
	for ; p != nil && p.data != data; i++ {
		p = p.next
	}
	if p == nil {
		return 0, errors.New("the element is not in the linklist")
	}
	return i, nil
}

// 单链表插入：在第i∈[1,n]个结点前插入元素为d的新结点，i可以插入到表长加一的地方
// 插入到第i个位置要找的是第i-1个结点，因为是对第i-1个结点的next域做替换
// 插入操作要从头结点开始，因为首元结点可能不存在，也就是说是空链表
func (head *LinkNode) Insert(i int, d Elem) error {
	// 需要判断i是否在链表长度+1范围内，在p == nil时还未找到就不在长度内，再判断此时的j == i-1如果成立说明是尾插
	// 换种思路，这种思路的代码明显更简洁
	p := head // p从头结点开始是考虑到头插
	j := 0
	for p != nil && j < i-1 { // 此处本来是j < i就能找到下一个，但是循环体中是 p = p.next 所以相当于是要找上上一个，也就是j < i-1了
		p = p.next
		j++ // 写在此处便于理解
	}
	if p == nil || j > i-1 { // p == nil处理i大于length+1的情况，j > i-1处理i等于0和小于0的情况，所以i可以是int
		return errors.New("i is not in the length+1 range")
	}
	q := &LinkNode{d, nil}
	q.next = p.next
	p.next = q
	return nil
}

// 单链表删除：删除第i∈[1,n]个结点，返回删除的结点
// 删除第i个结点，要找到第i-1个结点，将i-1结点的next指向第i个结点的next
// i有可能不在length范围内，
func (head *LinkNode) Delete(i int) (*LinkNode, error) {
	p := head // 指向头结点
	j := 0
	for p.next != nil && j < i-1 {
		p = p.next
		j++
	}

	if p.next == nil || j > i-1 {
		return nil, errors.New("i is not in the length range")
	}
	q := p.next
	p.next = p.next.next
	return q, nil
}

// 阶段小结 在循环体中是p = p.next时
// 判断i小于1 j>i-1
// 判断i大于length p.next == nil

// 单链表的建立：头插法
// 将要插入的n个元素按照从后往前的顺序依次插入到头结点之后的位置就是头插法
// 对一个只有头指针和头结点的空链表进行操作
// func (head *LinkNode) CreateLinkH(n int, v Elem) {
// 	// 依次将数据插入到头结点的位置
// 	// 学会接受i--
// 	for i := n; i > 0; i-- {
// 		var p *LinkNode
// 		p.data = v
// 		// fmt.Scanf("%d", p.data)
// 		p.next = head.next
// 		head.next = p
// 	}
// }

func (head *LinkNode) CreateLinkH(elems ...Elem) {
	// 依次将数据插入到头结点的位置
	// 学会接受i--
	for i := len(elems) - 1; i >= 0; i-- {
		p := &LinkNode{0, nil}
		p.data = elems[i]
		p.next = head.next
		head.next = p
	}
}

// 单链表的建立：尾插法
// 将要插入的n个元素按照从前往后的顺序依次插入到尾结点，使用一个尾指针指向尾结点
// 对一个只有头指针和头结点的空链表进行操作
// func (head *LinkNode) CreateLinkT(n int, v Elem) {
// 	// 依次将数据插入到尾结点的位置
// 	tail := head
// 	for i := 0; i < n; i++ {
// 		var p *LinkNode
// 		p.data = v
// 		// fmt.Scanf("%d", p.data)
// 		p.next = nil
// 		tail.next = p
// 		tail = p
// 	}
// }

func (head *LinkNode) CreateLinkT(elems ...Elem) {
	// 依次将数据插入到尾结点的位置
	tail := head
	for i := 0; i < len(elems); i++ {
		p := &LinkNode{0, nil}
		p.data = elems[i]
		p.next = nil
		tail.next = p
		tail = p
	}
}

// 打印链表
func (head *LinkNode) PrintLink() {
	// 从首元结点开始打印，打印到p.next=nil为止
	p := head.next

	for i := 1; p != nil; i++ {
		fmt.Printf("#%d p.data: %v    p.next: %p\n", i, p.data, p.next)
		p = p.next
	}
}
