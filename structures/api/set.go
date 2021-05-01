package api

type Set interface {
	StructureSize
	Add(val EqualHashRule)
	Contains(val EqualHashRule) bool
	Remove(val EqualHashRule) EqualHashRule
	ToArray() []EqualHashRule
}
