package selection

import (
	"fmt"

	"guochao.com/datastruct/08_sort/define"
)

// 堆排序 O(nlogn) 不稳定
// 从小到大，小根堆排序过程
// 如果将每次最小的根结点依次取出排成序列为从小到大排序
// l[1], l[i] = l[i], l[1] 如果将根（1号位置的值）与尾部的值进行交换，而不是直接访问取出，则形成的堆序列就是从大到小排序

func heapAdjust(l []define.RecordType, s int, n int) {
	// 已知l∈[s,m]中记录的关键字除l[s]外均满足堆的定义，本函数调整l[s]关键字，使l∈[s,m]成为一个小根堆
	// 记录当前s的值
	k := l[s]
	// 沿key较小的孩子结点向下筛选
	for j := 2 * s; j <= n; j *= 2 {
		// j在范围内，且选取较小的一个
		if j < n && l[j].Key > l[j+1].Key {
			j++
		}
		if k.Key <= l[j].Key {
			break
		}
		// 将较小的值加入到s位置上
		l[s] = l[j]
		s = j
	}
	// 将较大的值插入到最后一个被置换的地方
	l[s] = k
}

// 堆排序前置知识：1.在按层次编号的完全二叉树中第i个叶子结点的父节点一定是i/2，2.第j个结点的左孩子是2j、右孩子是2j+1，3.当第i个结点符合2i>n时一定是叶子结点
// 堆排序是将顺序表按照上述性质2构造为大根堆或者小根堆（根节点一定大于或小于孩子结点）之后，取根节点然后将最后一个值放在根节点位置，对其进行向下调整使其符合大小根堆得特征，反复上述过程，得到排序结果
func HeapSort(l []define.RecordType) {
	len := len(l)
	// 0号位置不存值
	n := len - 1
	// 将l创建为初始堆
	// 因为叶子结点符合大小根堆得定义，所以从第一个非叶子结点开始调整
	for i := n / 2; i >= 1; i-- {
		heapAdjust(l, i, n)
	}
	// 进行n-1趟排序(堆排序)
	res := ""
	for i := n; i >= 1; i-- {
		// 将根与最后一个元素交换
		res = fmt.Sprintf("%s->%d", res, l[1].Key)
		// 此步交换是关键，如果这样交换，最后形成的堆序列就是从大到小，而取出的每个根节点依然是从小到大
		l[1], l[i] = l[i], l[1]
		heapAdjust(l, 1, i-1)
	}
	fmt.Println(res)
}
