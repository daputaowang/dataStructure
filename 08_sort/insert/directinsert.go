package insert

import "guochao.com/datastruct/08_sort/define"

// 直接插入排序，稳定
// 因为排序不会引起切片扩容，所以不用返回

// 从第二个值开始起依次与前面的有序序列做比较，直到找到当前值合适的位置
func DirectInsertSort(l []define.RecordType) {
	len := len(l)
	// l[i]表示当前要与前面有序序列比较的值，j用来指向每次比较的位置
	var i, j int
	// 外部循环控制从第二个元素开始起进行比较然后插入
	for i = 2; i < len; i++ {
		// 判断当前值是否小于前一个，如果不小于则不需要进一步比较与移动
		if l[i].Key < l[i-1].Key {
			// 设置哨兵
			l[0] = l[i]
			// 从后向前比较
			for j = i - 1; l[0].Key < l[j].Key; j-- {
				// 移动
				l[j+1] = l[j]
			}
			// 找到插入位置的前一个位置，插入
			l[j+1] = l[0]
		}
	}
}

func DirectInsertSort1(nums []int) {
	if len(nums) == 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		// nums[j+1] < nums[j] 规定了小于才替换，相等不是小于所以不用替换
		for j := i - 1; j > -1 && nums[j+1] < nums[j]; j-- {
			nums[j+1], nums[j] = nums[j], nums[j+1]
		}
	}
}
