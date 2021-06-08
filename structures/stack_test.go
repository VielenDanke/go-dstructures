package structures

import "github.com/vielendanke/go-dstructures/structures/api"

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
