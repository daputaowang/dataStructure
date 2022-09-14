package linklist

import (
	"testing"
)

func TestMergeListLinkTola(t *testing.T) {
	// 引入表组测试
	testCases := []struct {
		name      string
		la        *LinkNode
		lb        *LinkNode
		laValue   []Elem
		lbValue   []Elem
		expcValue []Elem
	}{
		{name: "laOver", la: NewLinkNode(), lb: NewLinkNode(), laValue: []Elem{1, 7, 8}, lbValue: []Elem{2, 4, 6, 8, 10, 11}, expcValue: []Elem{1, 2, 4, 6, 7, 8, 8, 10, 11}},
		{name: "laNil", la: NewLinkNode(), lb: NewLinkNode(), laValue: []Elem{}, lbValue: []Elem{1}, expcValue: []Elem{1}},
		{name: "lbOver", la: NewLinkNode(), lb: NewLinkNode(), laValue: []Elem{2, 7, 8, 10}, lbValue: []Elem{1, 2, 5}, expcValue: []Elem{1, 2, 2, 5, 7, 8, 10}},
		{name: "lbNil", la: NewLinkNode(), lb: NewLinkNode(), laValue: []Elem{1}, lbValue: []Elem{}, expcValue: []Elem{1}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.la.CreateLinkT(testCase.laValue...)
			testCase.lb.CreateLinkT(testCase.lbValue...)
			MergeListLinkTola(testCase.la, testCase.lb)
			p := testCase.la.next
			for i := 0; p != nil; i++ {
				if p.data != testCase.expcValue[i] {
					t.Fatalf("received elem[%d] %d, expect %d", i, p.data, testCase.expcValue[i])
				}
				p = p.next
			}
			t.Log(testCase.name)
			testCase.la.PrintLink()
		})
	}

}

func TestLinkList(t *testing.T) {
	head := NewLinkNode()
	t.Logf("%v\n", &head)
	t.Logf("head.data:%d,head.next:%p\n", head.data, head.next)

	isEmpty := head.ListEmpty()
	t.Log("isEmpty", isEmpty)

	head.Insert(1, 2)
	head.Insert(2, 5)
	head.Insert(1, 1)
	head.Insert(2, 3)
	err1 := head.Insert(0, 1)
	err2 := head.Insert(6, 3)
	t.Log(err1, err2)
	head.PrintLink()

	length1 := head.LinkListLen()
	if length1 != 4 {
		t.Fatalf("received length %d, expect 4", length1)
	}

	loc, err3 := head.LocateElemIDX(10)
	if err3 != nil {
		t.Log(err3)
	}

	if loc != 0 {
		t.Fatalf("received elem %d, expect %d", loc, 0)
	}

	e, err4 := head.GetElem(10)

	if err4 != nil {
		t.Log(err4)
	}

	if e != 0 {
		t.Fatalf("received elem %d, expect %d", e, 0)
	}

	expcValue := []Elem{1, 3, 2, 5}
	p := head.next

	for i := 1; p != nil; i++ {
		if p.data != expcValue[i-1] {
			t.Fatalf("received elem[%d] %d, expect %d", i, p.data, expcValue[i-1])
		}
		p = p.next
	}

	h1, err5 := head.Delete(1)
	if err5 != nil {
		t.Log(err5)
	}

	if h1.data != 1 {
		t.Fatalf("received elem %d, expect %d", h1.data, 1)
	}

	head.PrintLink()
}
