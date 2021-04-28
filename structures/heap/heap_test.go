package heap

import (
	"os"
	"testing"
)

func prepareBinaryHeap() *binaryHeap {
	bh := NewBinaryHeap()
	for i := 1; i < 21; i++ {
		bh.Insert(i)
	}
	return bh
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
