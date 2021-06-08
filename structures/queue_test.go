package structures

import (
	"github.com/vielendanke/go-dstructures/structures/api"
)

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
