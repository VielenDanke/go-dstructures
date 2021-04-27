package queue

import (
	"fmt"
	"github.com/vielendanke/go-dstructures/structures/api"
)

type arrayQueue struct {
	elements []api.EqualHashRule
}

func NewArrayQueue() api.Queue {
	return &arrayQueue{}
}

func (aq *arrayQueue) Enqueue(val api.EqualHashRule) {
	aq.elements = append(aq.elements, val)
}

func (aq *arrayQueue) Dequeue() (api.EqualHashRule, bool) {
	if len(aq.elements) == 0 {
		return nil, false
	}
	defer aq.removeElement()
	return aq.elements[0], true
}

func (aq *arrayQueue) Size() int {
	return len(aq.elements)
}

func (aq *arrayQueue) String() string {
	return fmt.Sprintf("%v", aq.elements)
}

func (aq *arrayQueue) removeElement() {
	aq.elements = aq.elements[1:len(aq.elements)]
}
