package queue

import (
	"fmt"
	"testing"
)

func TestSqQueue(t *testing.T) {
	q := NewSqQueue(5)
	// q.Print()
	t.Logf("%v", q.s[:5])
	len1 := q.QueueLength()
	t.Log(len1)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	err1 := q.EnQueue(5)
	res1, _ := q.DeQueue()
	q.EnQueue(6)
	res2, _ := q.DeQueue()
	q.EnQueue(7)
	res3, _ := q.GetHead()
	fmt.Println(q.front, q.rear, err1, res1, res2, res3, q.QueueLength())
	q.Print()

}
