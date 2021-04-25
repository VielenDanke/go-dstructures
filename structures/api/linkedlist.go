package api

type LinkedList interface {
	StructureSize
	Stack
	Shift() (interface{}, bool)
	Unshift(val interface{})
	Get(idx int) (interface{}, bool)
	Set(idx int, val interface{}) bool
	Insert(idx int, val interface{}) bool
	Remove(idx int) (interface{}, bool)
	Reverse()
}
