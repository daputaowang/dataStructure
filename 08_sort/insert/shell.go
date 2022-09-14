package insert

import (
	"fmt"
	"math"

	"guochao.com/datastruct/08_sort/define"
)

// 希尔排序
// 不稳定（相同的值可能会出现在不同步长插入排序的过程中）

// 根据表长获取(2^k)-1规则的增量序列
func getIncrQ(len int) []int {
	// 因为数据元素是从1号位置开始计数的，所以len的值是实际长度减一的值
	// 先求出(2^k)-1的值最接近表长的k，而根据对数公式 (2^k)-1 = len-1 得 k = log2(len)。在此基础上减1的到一个k的最大值
	k := int(math.Floor(math.Log2(float64(len)))) - 1
	incrQ := make([]int, k)
	for i := 0; i < k; i++ {
		// incrQ[i] = int(math.Pow(2, float64(k-i))) - 1
		incrQ[i] = 1<<(k-i) - 1 // 相当于1*2^(k-i)
	}
	fmt.Println("增量序列为:", incrQ)
	return incrQ
}

// 按照步长进行一轮插入排序
func shellInsert(l []define.RecordType, s int) {
	len := len(l)
	// 外层循环控制当前比较的元素
	for i := s + 1; i < len; i++ {
		// 写入哨兵（此处的作用为记录元素）
		l[0] = l[i]
		var j int
		// 内存循环控制从当前元素位置按照步长往回插入排序并进行比较
		for j = i - s; j > 0 && l[0].Key < l[j].Key; j -= s {
			// 移动
			l[j+s] = l[j]
		}
		l[j+s] = l[0]
	}
}

// 希尔排序可以理解为多次进行步长逐渐缩小的插入排序
// 其中步长缩量之后最好是互质的，这个步长序列称为增量序列
func ShellSort(l []define.RecordType) {
	// 求增量序列按照(2^k)-1规则来取
	len := len(l)
	incrQ := getIncrQ(len)
	// 依次取出步长进行排序
	for s := range incrQ {
		shellInsert(l, s)
	}
}

// 简单希尔排序
// 使用希尔提出希尔排序时使用的增量序列step: len/2, len/2/2, len/2/2/2, ... 1
func ShellSortS(l []define.RecordType) {
	len := len(l)
	step := (len - 1) / 2
	for step > 0 {
		// 外层循环控制将要被插入的元素
		for i := step + 1; i < len; i++ {
			l[0] = l[i]
			var j int
			for j = i - step; j > 0 && l[0].Key < l[j].Key; j -= step {
				l[j+step] = l[j]
			}
			l[j+step] = l[0]
		}
		step = step / 2
	}
}

func ShellSortS1(nums []int) {
	step := len(nums) / 2
	// 最外层for循环控制步长
	for step > 0 {
		for i := step; i < len(nums); i += step {
			for j := i - step; j > -1 && nums[j+step] < nums[j]; j -= step {
				nums[j+step], nums[j] = nums[j], nums[j+step]
			}
		}
		step /= 2
	}
}
