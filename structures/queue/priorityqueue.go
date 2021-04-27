package queue

import (
	"fmt"
	"github.com/vielendanke/go-dstructures/structures/api"
	"math"
)

type priorityQueue struct {
	elements     []api.EqualHashRule
	calcPriority func(elem interface{}) int
}

func NewPriorityQueue(calcPriority func(elem interface{}) int) api.Queue {
	return &priorityQueue{calcPriority: calcPriority}
}

func (p *priorityQueue) Enqueue(val api.EqualHashRule) {
	p.elements = append(p.elements, val)
	p.enqueueWithPriority()
}

func (p *priorityQueue) Dequeue() (api.EqualHashRule, bool) {
	if p.Size() == 0 {
		return nil, false
	}
	max := p.elements[0]
	end := p.elements[p.Size()-1]
	p.elements = p.elements[:p.Size()-1]
	if p.Size() > 0 {
		p.elements[0] = end
		p.sinkDownWithPriority()
	}
	return max, true
}

func (p *priorityQueue) Size() int {
	return len(p.elements)
}

func (p *priorityQueue) String() string {
	return fmt.Sprintf("%v", p.elements)
}

func (p *priorityQueue) enqueueWithPriority() {
	idx := p.Size() - 1
	element := p.elements[idx]
	for idx > 0 {
		parentIdx := int(math.Floor(float64(idx-1) / 2))
		parent := p.elements[parentIdx]
		if p.calcPriority(element) >= p.calcPriority(parent) {
			break
		}
		p.elements[parentIdx] = element
		p.elements[idx] = parent
		idx = parentIdx
	}
}

func (p *priorityQueue) sinkDownWithPriority() {
	idx := 0
	length := p.Size()
	element := p.elements[0]
	for {
		leftChildIdx := 2*idx + 1
		rightChildIdx := 2*idx + 2
		var leftChild, rightChild interface{}
		swap := -1
		if leftChildIdx < length {
			leftChild = p.elements[leftChildIdx]
			if p.calcPriority(leftChild) < p.calcPriority(element) {
				swap = leftChildIdx
			}
		}
		if rightChildIdx < length {
			rightChild = p.elements[rightChildIdx]
			if (swap == -1 && p.calcPriority(rightChild) < p.calcPriority(element)) ||
				(swap != -1 && p.calcPriority(rightChild) < p.calcPriority(leftChild)) {
				swap = rightChildIdx
			}
		}
		if swap == -1 {
			break
		}
		p.elements[idx] = p.elements[swap]
		p.elements[swap] = element
		idx = swap
	}
}
