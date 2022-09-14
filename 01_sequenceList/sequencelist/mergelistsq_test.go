package sequencelist

import (
	"testing"
)

// func TestMergeListSq(t *testing.T) {
// 	// 前面的先结束的情况
// 	la := SqList{
// 		elem:   []int{1, 7, 8},
// 		length: 3,
// 	}
// 	lb := SqList{
// 		elem:   []int{2, 4, 6, 8, 10, 11},
// 		length: 6,
// 	}
// 	var lc SqList
// 	MergeListSq(la, lb, &lc)
// 	expc := []int{1, 2, 4, 6, 7, 8, 8, 10, 11}
// 	explc := SqList{elem: expc, length: 9}
// 	compare(lc, explc, t)

// 	// 前面为nil的情况
// 	la2 := SqList{
// 		elem:   []int{},
// 		length: 0,
// 	}
// 	lb2 := SqList{
// 		elem:   []int{1},
// 		length: 1,
// 	}
// 	var lc2 SqList
// 	MergeListSq(la2, lb2, &lc2)
// 	expc2 := []int{1}
// 	explc2 := SqList{elem: expc2, length: 1}
// 	compare(lc2, explc2, t)

// 	// 后面的先结束的情况
// 	la1 := SqList{
// 		elem:   []int{2, 7, 8, 10},
// 		length: 4,
// 	}
// 	lb1 := SqList{
// 		elem:   []int{1, 2, 5},
// 		length: 3,
// 	}
// 	var lc1 SqList
// 	MergeListSq(la1, lb1, &lc1)
// 	expc1 := []int{1, 2, 2, 5, 7, 8, 10}
// 	explc1 := SqList{elem: expc1, length: 7}
// 	compare(lc1, explc1, t)

// 	// 后面为nil的情况
// 	la3 := SqList{
// 		elem:   []int{1},
// 		length: 1,
// 	}
// 	lb3 := SqList{
// 		elem:   []int{},
// 		length: 0,
// 	}
// 	var lc3 SqList
// 	MergeListSq(la3, lb3, &lc3)
// 	expc3 := []int{1}
// 	explc3 := SqList{elem: expc3, length: 1}
// 	compare(lc3, explc3, t)
// }

// func compare(res SqList, exp SqList, t *testing.T) {
// 	if exp.length != res.length {
// 		t.Fatalf("received length %d, expect %d", res.length, exp.length)
// 	}

// 	for i, v := range res.elem {
// 		if v != exp.elem[i] {
// 			t.Fatalf("received elem[%d] %d, expect %d", i, v, exp.elem[i])
// 		}
// 	}
// }

func TestMergeListSq(t *testing.T) {
	// 引入表组测试
	testCases := []struct {
		name string
		la   SqList
		lb   SqList
		expc SqList
	}{
		{name: "laOver", la: SqList{[]int{1, 7, 8}, 3}, lb: SqList{[]int{2, 4, 6, 8, 10, 11}, 6}, expc: SqList{[]int{1, 2, 4, 6, 7, 8, 8, 10, 11}, 9}},
		{name: "laNil", la: SqList{[]int{}, 0}, lb: SqList{[]int{1}, 1}, expc: SqList{[]int{1}, 1}},
		{name: "lbOver", la: SqList{[]int{2, 7, 8, 10}, 4}, lb: SqList{[]int{1, 2, 5}, 3}, expc: SqList{[]int{1, 2, 2, 5, 7, 8, 10}, 7}},
		{name: "lbNil", la: SqList{[]int{1}, 1}, lb: SqList{[]int{1}, 0}, expc: SqList{[]int{1}, 1}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var lc SqList
			MergeListSq(testCase.la, testCase.lb, &lc)

			if testCase.expc.length != lc.length {
				t.Fatalf("received length %d, expect %d", lc.length, testCase.expc.length)
			}

			for i, v := range lc.elem {
				if v != testCase.expc.elem[i] {
					t.Fatalf("received elem[%d] %d, expect %d", i, v, testCase.expc.elem[i])
				}
			}
		})
	}
}
