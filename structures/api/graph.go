package api

type Graph interface {
	StructureSize
	AddVertex(val EqualHashRule) bool
	AddEdge(fVertex EqualHashRule, sVertex EqualHashRule) bool
	RemoveEdge(fVertex EqualHashRule, sVertex EqualHashRule) bool
	RemoveVertex(fVertex EqualHashRule) bool
	Contains(vertex EqualHashRule) bool
	GetEdges(vertex EqualHashRule) []EqualHashRule
	ToArray() []EqualHashRule
}
