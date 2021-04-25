package api

type Queue interface {
	StructureSize
	Enqueue(val interface{})
	Dequeue() (interface{}, bool)
}
