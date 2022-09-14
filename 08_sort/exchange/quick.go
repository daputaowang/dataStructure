package exchange

import "guochao.com/datastruct/08_sort/define"

// 快速排序 O(nlogn) 不稳定
// 随机选取一个中轴（pivot）值，将大于等于中轴值的都放在中轴值的左边，将小于等于中轴值的放在右边
// 对左右两边重复上述过程
func QuickSort(l []define.RecordType, Low int, High int) {
	// low大于等于high时递归退出
	// 当序列的长度为1时，low = high，此判断不成立退出该层递归
	if Low < High {
		// 保留每次递归时传入的区间值，以便于下次递归时划分新的区间
		low, high := Low, High
		// 将第一个元素记作中轴值，存入0位置(中轴值可以随机选取)
		l[0] = l[low]
		pivot := l[low].Key
		// 外层循环控制将数据根据中轴值交换到左右两边
		for low < high {
			// 判断high处的值是否大于等于pivot的值，若是high--，左移
			for low < high && l[high].Key >= pivot {
				high--
			}
			// 若不是，将high处的值放在low处，high处的值就空出来了
			l[low] = l[high]

			// 判断low处的值是否小于pivot的值，若是low++，右移
			for low < high && l[low].Key < pivot {
				low++
			}
			// 若不是，将low处的值，放在high处
			l[high] = l[low]
		}
		// 循环结束low=high，将中轴值放在此处
		l[low] = l[0]
		// 对小于pivot的部分进行递归排序
		QuickSort(l, Low, low-1)
		// 对大于pivot的部分进行递归排序
		QuickSort(l, low+1, High)
	}
}

func QuickSort1(nums []int, left, right int) {
	// 分割到只剩一个元素left == right, 分割到一个元素都没有left > right
	if left >= right {
		return
	}
	// partition分区过程，把第一个元素作为中轴pivot，i探子从左到右探索，>pivot跳过，<=pivot的跟大序列的第一个元素交换，从而把元素挪到小序列的最后
	// j用来标记小序列队尾，当小序列长度为0时，队尾就在-1的位置，小序列的-1的位置就是left的位置。
	// i探子探索结束之后，将小序列队尾j位置的元素与left位置的元素做交换，这样pivot左右就都是小于等于pivot的元素，右侧都是大于pivot的元素。
	pivot := nums[left]
	j := left
	for i := left + 1; i <= right; i++ {
		if nums[i] <= pivot {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[left], nums[j] = nums[j], nums[left]
	QuickSort1(nums, left, j-1)
	QuickSort1(nums, j+1, right)
}
