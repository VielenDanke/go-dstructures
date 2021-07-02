package api

type Queue interface {
	Structure
	Enqueue(val EqualHashRule)
	Dequeue() (EqualHashRule, bool)
}
