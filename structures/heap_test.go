package structures

func prepareBinaryHeap() *binaryHeap {
	bh := NewBinaryHeap()
	for i := 1; i < 21; i++ {
		bh.Insert(i)
	}
	return bh
}
