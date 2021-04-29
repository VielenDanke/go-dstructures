package api

type Graph interface {
	StructureSize
	AddVertex(val EqualHashRule) bool
	AddEdge(fVertex EqualHashRule, sVertex EqualHashRule) bool
	RemoveEdge(fVertex EqualHashRule, sVertex EqualHashRule) bool
	RemoveVertex(fVertex EqualHashRule) bool
	ToArray() []EqualHashRule
}
