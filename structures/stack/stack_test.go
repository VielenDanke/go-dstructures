package stack

import "github.com/vielendanke/go-dstructures/structures/api"

type cInt int

func (c cInt) Equal(p interface{}) bool {
	return c.Hash() == p.(cInt).Hash()
}

func (c cInt) Hash() int {
	return int(c)
}

func prepareArrayStack() api.Stack {
	return fillStack(NewArrayStack())
}

func prepareLinkedStack() api.Stack {
	return fillStack(NewLinkedStack())
}

func fillStack(s api.Stack) api.Stack {
	for i := 1; i < 21; i++ {
		s.Push(cInt(i))
	}
	return s
}
