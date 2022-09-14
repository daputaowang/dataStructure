package exchange

import (
	"guochao.com/datastruct/08_sort/define"
)

// 交换排序中的冒泡排序 稳定

// 冒泡排序思想
// 从头到尾依次将前后两个值比较，将大的放在后面，比较一趟就可以将最大的沉在最后，将小的阶段性向前移动
// 一共需要比较n-1趟(表长减一趟)，第m趟比较n-m次
// 表的0号位置为空
func BubbleSort(l []define.RecordType) {
	len := len(l)
	n := len - 1
	// 外层循环控制冒泡趟数
	for m := 1; m <= n-1; m++ {
		// 内层循环控制每趟的比较
		for i := 1; i <= n-m; i++ {
			if l[i].Key > l[i+1].Key {
				l[i], l[i+1] = l[i+1], l[i]
			}
		}
	}
}

func BubbleFlagSort(l []define.RecordType) {
	len := len(l)
	n := len - 1
	flag := 1
	// 外层循环控制冒泡趟数
	for m := 1; m <= n-1 && flag == 1; m++ {
		flag = 0
		// 内层循环控制每趟的比较
		for i := 1; i <= n-m; i++ {
			if l[i].Key > l[i+1].Key {
				flag = 1
				l[i], l[i+1] = l[i+1], l[i]
			}
		}
	}
}

func BubbleSort1(nums []int) {
	// 一共需要比较表长减1次，因为最后一趟时是前两个比较。
	m := len(nums) - 1
	flag := 1
	for i := 0; i < m && flag == 1; i++ {
		flag = 0
		for j := 0; j < m-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = 1
			}
		}
	}
}
