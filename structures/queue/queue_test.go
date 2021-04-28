package queue

import (
	"github.com/vielendanke/go-dstructures/structures/api"
	"os"
	"testing"
)

type cInt int

func (c cInt) Equal(p interface{}) bool {
	return c.Hash() == p.(cInt).Hash()
}

func (c cInt) Hash() int {
	return int(c)
}

var priorityCalc = func(elem interface{}) int {
	iElem := elem.(cInt)
	return int(iElem)
}

func preparePriorityQueue() api.Queue {
	return fillQueue(NewPriorityQueue(priorityCalc))
}

func prepareArrayQueue() api.Queue {
	return fillQueue(NewArrayQueue())
}

func fillQueue(q api.Queue) api.Queue {
	for i := 21; i > 0; i-- {
		q.Enqueue(cInt(i))
	}
	return q
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
