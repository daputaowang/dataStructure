package linklist

import "fmt"

// 合并顺序表la与lb到lc中
func MergeListLink(la *LinkNode, lb *LinkNode, lc *LinkNode) {
	// 从首元结点开始比较，较小的放入lc中
	pa, pb, pc := la.next, lb.next, lc

	// 得到la与lb的首元结点的指针之后，就将la与lb释放
	la, lb = nil, nil

	// pa != nil 处理la先结束或者la为nil的情况，pb != nil 处理lb先结束或者lb为nil的情况
	for pa != nil && pb != nil {
		// 相等的情况pa中的值优先放入c中，所以此处使用小于等于
		// 先将pa指向的结点连接到pc指向的结点上，然后将pc移动到这个新的结点，然后将pa向后移动
		if pa.data <= pb.data {
			pc.next = pa
			pc = pa
			pa = pa.next
		} else {
			// 若不小于等于就是pa指向的值较小，则将pb指向结点连接到pc指向的结点上，然后将pc移动到这个新的结点，然后将pb向后移动
			pc.next = pb
			pc = pb
			pb = pb.next
		}
	}

	// la先结束就把lb剩下的直接链在lc中，lb先结束就把la剩下的直接链在lc中
	// 处理lb先结束la还有剩余的情况
	if pa != nil {
		pc.next = pa
	}

	// 处理la先结束lb还有剩余的情况
	if pb != nil {
		pc.next = pb
	}
}

// 合并顺序表la与lb到la中
func MergeListLinkTola(la *LinkNode, lb *LinkNode) {
	// 从首元结点开始比较，较小的放入lc中
	// pa与pb指向la与lb的首元结点，pc指向la的head，将后面的结点链到la
	pa, pb, pc := la.next, lb.next, la

	// 得到la与lb的首元结点的指针之后，就将la与lb释放
	lb = nil

	// pa != nil 处理la先结束或者la为nil的情况，pb != nil 处理lb先结束或者lb为nil的情况
	for pa != nil && pb != nil {
		// 相等的情况pa中的值优先放入c中，所以此处使用小于等于
		// 先将pa指向的结点连接到pc指向的结点上，然后将pc移动到这个新的结点，然后将pa向后移动
		if pa.data <= pb.data {
			pc.next = pa
			pc = pa
			pa = pa.next
		} else {
			// 若不小于等于就是pa指向的值较小，则将pb指向结点连接到pc指向的结点上，然后将pc移动到这个新的结点，然后将pb向后移动
			pc.next = pb
			pc = pb
			pb = pb.next
		}
	}

	// la先结束就把lb剩下的直接链在lc中，lb先结束就把la剩下的直接链在lc中
	// 处理lb先结束la还有剩余的情况
	if pa != nil {
		pc.next = pa
	}

	// 处理la先结束lb还有剩余的情况
	if pb != nil {
		pc.next = pb
	}
}

func MergeListLinkExec() {
	la := NewLinkNode()
	la.CreateLinkH(1, 7, 8)
	fmt.Println("la:")
	la.PrintLink()

	lb := NewLinkNode()
	lb.CreateLinkH(2, 4, 6, 8, 10, 11)
	fmt.Println("lb:")
	lb.PrintLink()

	lc := NewLinkNode()

	MergeListLink(la, lb, lc)
	fmt.Println("lc:")
	lc.PrintLink()

	// 两个链表中的结点不能是同一个结点，否则在p.next的过程中会出现环，之后的过程中就会陷入死循环了
	ld := NewLinkNode()
	ld.CreateLinkH(1, 7, 8)
	fmt.Println("ld:")
	ld.PrintLink()

	le := NewLinkNode()
	le.CreateLinkH(2, 4, 6, 8, 10, 11)
	fmt.Println("le:")
	le.PrintLink()
	MergeListLinkTola(ld, le)
	fmt.Println("ld:")
	ld.PrintLink()
}
