package api

type Set interface {
	StructureSize
	Add(val EqualHashRule)
	Contains(val EqualHashRule) bool
}
