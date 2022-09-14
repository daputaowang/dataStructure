package queue

import (
	"testing"
)

func TestLinkQueue(t *testing.T) {
	lq := NewLinkQueue()
	// lq.EnQueue(1)
	// lq.EnQueue(2)
	// lq.EnQueue(3)
	// res2, _ := lq.DeQueue()
	// res1, _ := lq.GetHead()
	// fmt.Println(res2, res1, lq.length)
	lq.CreateLinkQueueT(1, 2, 3)
	lq.DestroyQueue()
	lq.PrintLinkQueue()
}
