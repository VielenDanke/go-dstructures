package api

type Set interface {
	StructureSize
	Structure
	Add(val EqualHashRule)
	Remove(val EqualHashRule) EqualHashRule
}
