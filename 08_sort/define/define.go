package define

type RecordType struct {
	Key  int
	Info string
}

// 第一个位置不存值，用来做哨兵
// 排序时，均做从小到大排序
func NewList() []RecordType {
	return []RecordType{{0, "0"}, {10, "1"}, {5, "2"}, {3, "3"}, {16, "4"}, {7, "5"}, {32, "6"}, {83, "7"}, {23, "8"}, {54, "9"}, {29, "10"}}
}

func NewSlict() []int {
	return []int{0, 10, 5, 3, 16, 7, 32, 83, 23, 54, 29}
}
