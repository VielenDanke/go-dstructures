package structures

import (
	"fmt"

	"github.com/vielendanke/go-dstructures/structures/api"
)

type arrayStack struct {
	elements []api.EqualHashRule
}

func NewArrayStack() api.Stack {
	return &arrayStack{}
}

func (as *arrayStack) Push(val api.EqualHashRule) {
	as.elements = append(as.elements, val)
}

func (as *arrayStack) Pop() (val api.EqualHashRule, isFound bool) {
	if len(as.elements) == 0 {
		return
	}
	val = as.elements[len(as.elements)-1]
	defer as.removeElemAfterPop()
	return val, !isFound
}

func (as *arrayStack) Size() int {
	return len(as.elements)
}

func (as *arrayStack) String() string {
	return fmt.Sprintf("%v", as.elements)
}

func (as *arrayStack) removeElemAfterPop() {
	as.elements = as.elements[0 : len(as.elements)-1]
}
