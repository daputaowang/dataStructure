package stack

import "testing"

func TestSqStack(t *testing.T) {
	stp := NewStack(4)
	t.Logf("%v\n", stp)

	ok1 := stp.StackEmpty()
	if !ok1 {
		t.Errorf("received %t, expect true\n", ok1)
	}

	err1 := stp.Push(1)
	err2 := stp.Push(2)
	err3 := stp.Push(3)
	err4 := stp.Push(4)

	if err1 != nil && err2 != nil && err3 != nil && err4 != nil {
		t.Log(err1, err2, err3, err4)
	}

	stp.Print()

	stplength := stp.StackLength()
	if stplength != 4 {
		t.Errorf("received %d, expect 4\n", stplength)
	}

	err5 := stp.Push(5)
	t.Log(err5)
	if err5 == nil {
		t.Errorf("received %v, expect stack overflow\n", err5)
	}

	res1, err6 := stp.Pop()
	t.Logf("res %d, err %v\n", res1, err6)
	if res1 != 4 {
		t.Errorf("received %d, expect 4\n", res1)
	}

	stp.Print()

	res2, err7 := stp.Peek()
	t.Logf("res %d, err %v\n", res2, err7)
	if res2 != 3 {
		t.Errorf("received %d, expect 3\n", res2)
	}

	stp.Print()
}
