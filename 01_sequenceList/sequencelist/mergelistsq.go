package sequencelist

import "fmt"

// 有序表合并-顺序表实现
// 解题思路
// 1.记录要合并的两个顺序表的头尾位置
// 2.进行循环比较，谁的小谁先放入lc，保持la排序后的稳定性，所以la是小于等于lb
// 3.当la与lb中有一个顺序表的数据全部进入了lc则将另一个表的数据全部迁移到lc中
func MergeListSq(la SqList, lb SqList, lc *SqList) {
	lc.length = la.length + lb.length
	lc.elem = make([]int, lc.length)

	lowa, lowb := 0, 0
	lowc := 0
	// lowa <= int(higha)处理la先结束以及la为空的情况，lowb <= int(highb)处理lb先结束以及lb为空的情况
	for lowa <= int(la.length-1) && lowb <= int(lb.length-1) {
		if la.elem[lowa] <= lb.elem[lowb] {
			lc.elem[lowc] = la.elem[lowa]
			lowa++
		} else {
			lc.elem[lowc] = lb.elem[lowb]
			lowb++
		}
		lowc++
	}

	// 处理lb先结束la还有剩余的情况
	for lowa <= int(la.length-1) {
		lc.elem[lowc] = la.elem[lowa]
		lowa++
		lowc++
	}

	// 处理la先结束lb还有剩余的情况
	for lowb <= int(lb.length-1) {
		lc.elem[lowc] = lb.elem[lowb]
		lowb++
		lowc++
	}
}

// 有序表合并-顺序表实现，将lb表合并到la中，la有足够的空间
// 解题思路
// 1.记录两个有序表的首尾位置
// 2.进行循环比较，
func MergeListSq_Merge2La(la *SqList, lb SqList) {
	lowa, lowb := 0, 0
	higha := la.length - 1
	highb := lb.length - 1

	for lowa <= int(higha) && lowb <= int(highb) {
		if lb.elem[lowb] <= la.elem[lowa] {
			la.Insert(uint(lowa), lb.elem[lowb])
			lowb++
			lowa++
			higha++
		} else {
			lowa++
		}
	}

	for lowb <= int(highb) {
		la.Insert(uint(lowa), lb.elem[lowb])
		lowa++
		lowb++
		higha++
	}

	// la.length = higha + 1
}

// func merge(A []int, m int, B []int, n int)  {
//     // 获取AB的首尾，判断A的最后一个元素大还是B的最后一个元素大，将大的放入A的末端
//     higha, highb, index := m-1, n-1, len(A)-1
//     // higha >=0 处理A先结束以及A为空的情况，highb >= 0 处理B先结束以及B为空的情况
//     for higha >= 0 && highb >= 0 {
//         if B[highb] >= A[higha] {
//             A[index] = B[highb]
//             highb--
//         } else {
//             A[index] = A[higha]
//             higha--
//         }
//         index--
//     }

//     // 处理A先结束，B还未结束的情况
//     for highb >= 0 {
//         A[index] = B[highb]
//         index--
//         highb--
//         // 其他人的解法
//         // copy(A[:highb+1], B[:highb+1])
//     }
// }

func MergeListSqExec() {
	// 有序表合并-顺序表实现
	// 前面先结束的情况
	la := SqList{
		elem:   []int{1, 7, 8},
		length: 3,
	}
	lb := SqList{
		elem:   []int{2, 4, 6, 8, 10, 11},
		length: 6,
	}
	var lc SqList
	MergeListSq(la, lb, &lc)
	fmt.Println("lc:", lc)

	LA := SqList{
		elem:   []int{1, 7, 8, 0, 0, 0, 0, 0, 0},
		length: 3,
	}

	MergeListSq_Merge2La(&LA, lb)
	fmt.Println("LA:", LA)

	// 后面的先结束的情况
	la1 := SqList{
		elem:   []int{2, 7, 8, 10},
		length: 4,
	}
	lb1 := SqList{
		elem:   []int{1, 2, 5},
		length: 3,
	}
	var lc1 SqList
	MergeListSq(la1, lb1, &lc1)
	fmt.Println("lc1:", lc1)
}
