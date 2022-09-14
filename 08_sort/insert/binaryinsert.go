package insert

import "guochao.com/datastruct/08_sort/define"

// 二分插入排序/折半插入排序， 稳定

// 将直接插入排序中在有序序列中查找插入位置的过程改为二分查找
func BinaryInsertSort(l []define.RecordType) {
	len := len(l)
	// 外层循环控制将要被插入的元素
	for i := 2; i < len; i++ {
		// 将l[0]设置为哨兵
		l[0] = l[i]
		// 查找的区间是[1, i-1]
		low, high := 1, i-1
		// 循环条件low <= high
		for low <= high {
			mid := low + (high-low)/2
			// 用小于号先判断小于的左半区间可以保证找到的位置的稳定性
			if l[0].Key < l[mid].Key {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		// 循环退出时high的位置一定在要插入位置的前一个
		// 将high+1到i-1位置的值，向后移动一位
		for j := i - 1; j > high; j-- {
			l[j+1] = l[j]
		}
		l[high+1] = l[0]
	}
}

func BinaryInsertSort1(nums []int) {
	if len(nums) == 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		// nums[j+1] < nums[j] 规定了小于才替换，相等不是小于所以不用替换
		// 使用二分查找确定第i个元素插入的位置
		low, high := 0, i-1
		for low <= high {
			mid := low + (high-low)/2
			if nums[i] < nums[mid] {
				high = mid - 1
			} else { // 等价于else if nums[i] >= nums[mid]，目标值大于或等于中间值时，low向上移动，这样会使得当目标值在序列中出现连续多个时，找到high边界
				low = mid + 1
			}
		}
		// low的位置就是插入的位置
		for j := i - 1; j >= low; j-- {
			nums[j+1], nums[j] = nums[j], nums[j+1]
		}
	}
}
