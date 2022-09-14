package stack

import "testing"

func TestLinkStack(t *testing.T) {
	lsp := NewLinkStack()

	lsp.CreateLinkT(4, 5, 6)
	lsp.PrintLink()

	t.Logf("%v", lsp.top)
	lsp.Push(1)
	t.Logf("%d", lsp.top.data)
	lsp.Push(2)
	t.Logf("%d", lsp.top.data)
	lsp.Push(3)
	t.Logf("%d", lsp.top.data)
	lsp.Push(4)
	t.Logf("%d", lsp.top.data)
	lsp.PrintLink()

	res1, err1 := lsp.Pop()
	if err1 != nil {
		t.Log(err1)
	}
	if res1 != 4 {
		t.Errorf("Pop error:received %d，expect 4\n", res1)
	}

	res2 := lsp.Peek()
	t.Log(res2, lsp.top.data)
	if res2 != 3 {
		t.Errorf("Peek error:received %d，expect 3", res2)
	}
}
