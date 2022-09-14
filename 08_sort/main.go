package main

import (
	"fmt"

	"guochao.com/datastruct/08_sort/define"
	// "guochao.com/datastruct/08_sort/insert"
	"guochao.com/datastruct/08_sort/exchange"
)

func main() {
	// list := define.NewList()
	lists := define.NewSlict()
	// =======================插入排序=========================
	// 直接插入排序
	// insert.DirectInsertSort(list)
	// insert.DirectInsertSort1(lists)
	// fmt.Println(lists)

	// 二分插入排序
	// insert.BinaryInsertSort(list)
	// fmt.Println(list[1:])
	// insert.BinaryInsertSort1(lists)
	// fmt.Println(lists)

	// 希尔排序
	// insert.ShellSortS(list)
	// fmt.Println(list[1:])
	// insert.ShellSortS1(lists)
	// fmt.Println(lists)

	// =======================交换排序=========================
	// 冒泡排序
	// exchange.BubbleFlagSort(list)
	// fmt.Println(list[1:])
	// exchange.BubbleSort1(lists)
	// fmt.Println(lists)

	// 快速排序
	// exchange.QuickSort(list, 1, len(list)-1)
	// fmt.Println(list[1:])
	exchange.QuickSort1(lists, 0, len(lists)-1)
	fmt.Println(lists)

	// =======================选择排序=========================
	// 简单选择排序
	// selection.SelectSort(list)
	// fmt.Println(list[1:])

	// 堆排序
	// selection.HeapSort(list)
	// fmt.Println(list)

	// =======================归并排序=========================
	// alist := make([]int, 30)
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 30; i++ {
	// 	alist[i] = rand.Intn(100)
	// }
	// res := merge.MergeSort(alist)
	// fmt.Println(res)
}
