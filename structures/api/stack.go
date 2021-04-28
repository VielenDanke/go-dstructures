package api

type Stack interface {
	StructureSize
	Push(val EqualHashRule)
	Pop() (EqualHashRule, bool)
}
