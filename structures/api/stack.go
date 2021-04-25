package api

type Stack interface {
	StructureSize
	Push(val interface{})
	Pop() (interface{}, bool)
}
