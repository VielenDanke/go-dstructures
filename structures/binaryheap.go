package structures

import (
	"fmt"
	"math"
)

/*
root -> left (2n+1) right (2n+2)
*/

type binaryHeap struct {
	elements []int
}

func NewBinaryHeap() *binaryHeap {
	return &binaryHeap{}
}

func (bh *binaryHeap) Insert(val int) {
	bh.elements = append(bh.elements, val)
	if bh.Size() == 1 {
		return
	} else {
		bh.swap(len(bh.elements) - 1)
	}
}

func (bh *binaryHeap) RemoveMax() (int, bool) {
	if bh.Size() == 0 {
		return -1, false
	}
	startIdx := 0
	elem := bh.elements[startIdx]
	if bh.Size() == 1 {
		bh.elements = bh.elements[startIdx : bh.Size()-1]
	} else {
		bh.elements[startIdx], bh.elements[bh.Size()-1] = bh.elements[bh.Size()-1], bh.elements[startIdx]
		bh.elements = bh.elements[startIdx : bh.Size()-1]
		bh.sinkDown(startIdx)
	}
	return elem, true
}

func (bh *binaryHeap) Contains(val int) bool {
	if bh.Size() == 0 {
		return false
	}
	if bh.Size() == 1 {
		return bh.elements[0] == val
	} else {
		return bh.containsDown(0, val)
	}
}

func (bh *binaryHeap) containsDown(idx int, val int) bool {
	elem := bh.elements[idx]
	if val > elem {
		return false
	} else {
		if val == elem {
			return true
		}
		var leftElem, rightElem int
		left := 2*idx + 1
		right := 2*idx + 2
		if left < bh.Size() {
			leftElem = bh.elements[left]
		} else {
			leftElem = -1
		}
		if right < bh.Size() {
			rightElem = bh.elements[right]
		} else {
			rightElem = -1
		}
		if val > leftElem && val > rightElem {
			return false
		}
		return bh.containsDown(left, val) || bh.containsDown(right, val)
	}
}

func (bh *binaryHeap) sinkDown(idx int) {
	left := 2*idx + 1
	right := 2*idx + 2
	var leftElem, rightElem, maxIdx, maxNum int
	if left > bh.Size() && right > bh.Size() {
		return
	}
	if left < bh.Size() {
		leftElem = bh.elements[left]
	}
	if right < bh.Size() {
		rightElem = bh.elements[right]
	}
	if leftElem > rightElem {
		maxNum = leftElem
		maxIdx = left
	} else {
		maxNum = rightElem
		maxIdx = right
	}
	if bh.elements[idx] > maxNum {
		return
	}
	if maxNum > bh.elements[idx] {
		bh.elements[idx], bh.elements[maxIdx] = bh.elements[maxIdx], bh.elements[idx]
		idx = maxIdx
		bh.sinkDown(idx)
	}
}

func (bh *binaryHeap) swap(idx int) {
	parentIdx := int(math.Floor((float64(idx) - 1) / 2))
	if idx <= 0 {
		return
	}
	if bh.elements[idx] > bh.elements[parentIdx] {
		bh.elements[idx], bh.elements[parentIdx] = bh.elements[parentIdx], bh.elements[idx]
		idx = parentIdx
		bh.swap(idx)
	}
}

func (bh *binaryHeap) Size() int {
	return len(bh.elements)
}

func (bh *binaryHeap) String() string {
	return fmt.Sprintf("%v", bh.elements)
}
