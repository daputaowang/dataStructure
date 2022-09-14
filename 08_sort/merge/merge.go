package merge

// 归并排序 O(nlogn) 稳定

// 合并两个顺序表为一个
// 思想为，使用两个表第一个元素的指针指向两个表
// 比较两个指针指向的位置的值，将较小的放入结果集中，向后移动较小的所在的表的指针再次进行比较
// 直到某一个表中元素被取完为止
func merge(la, lb []int) []int {
	lowa, lowb := 0, 0
	lena, lenb := len(la), len(lb)

	res := make([]int, lena+lenb)
	lowr := 0
	// 同时从头到位挨个比较两个表
	for lowa < lena && lowb < lenb {
		// 比较两个表的元素中较小的值
		if la[lowa] <= lb[lowb] {
			res[lowr] = la[lowa]
			lowa++
		} else {
			res[lowr] = lb[lowb]
			lowb++
		}
		lowr++
	}

	// la还有剩余时
	for lowa < lena {
		res[lowr] = la[lowa]
		lowa++
		lowr++
	}

	// lb还有剩余时
	for lowb < lenb {
		res[lowr] = lb[lowb]
		lowb++
		lowr++
	}
	return res
}

// 归并排序的思想为：
// 先递归将待排序集合二分拆解，到集合的长度为1为止
// 对长度为1的集合按照拆解的顺序左右合并
func MergeSort(l []int) []int {
	// 递归拆分数组
	n := len(l)
	if n <= 1 {
		return l
	}
	mid := n / 2
	lowL := MergeSort(l[:mid])
	highL := MergeSort(l[mid:])
	res := merge(lowL, highL)
	return res
}
