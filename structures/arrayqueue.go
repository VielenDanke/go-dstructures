package structures

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

func (aq *arrayQueue) Contains(val api.EqualHashRule) bool {
	for _, v := range aq.elements {
		isEqual := v.Equal(val)
		if isEqual {
			return true
		}
	}
	return false
}

func (aq *arrayQueue) ToArray() []api.EqualHashRule {
	return aq.elements
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
