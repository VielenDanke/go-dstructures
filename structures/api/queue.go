package api

type Queue interface {
	StructureSize
	Enqueue(val EqualHashRule)
	Dequeue() (EqualHashRule, bool)
}
