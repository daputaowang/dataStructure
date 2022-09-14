package selection

import "guochao.com/datastruct/08_sort/define"

// 简单选择排序
// 对长度为n的序列，进行n-1次比较，直到找到最小值，放在最前面，第一趟一共比较n-1次
// 保证第一个是最小的之后，就开始从第二个开始的序列重复上述过程，直到全部排完。一共比较n-1趟
func SelectSort(l []define.RecordType) {
	len := len(l)
	n := len - 1 // 0号位置不存值
	// 外层循环控制趟数
	for i := 1; i <= n-1; i++ {
		// 将第i个值作为较小值，记录最小值的位置，当向后比较的过程中出现更小的值，更新此位置
		k := i
		// 内层循环寻找最小的值
		for j := i; j < len; j++ {
			if l[k].Key > l[j].Key {
				k = j
			}
		}
		if k != i {
			l[i], l[k] = l[k], l[i]
		}
	}
}
