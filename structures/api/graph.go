package api

type Graph interface {
	StructureSize
	Structure
	AddVertex(val EqualHashRule)
	AddEdge(fVertex EqualHashRule, sVertex EqualHashRule, weight int64) bool
	RemoveEdge(fVertex EqualHashRule, sVertex EqualHashRule) bool
	RemoveVertex(fVertex EqualHashRule) bool
	GetEdges(vertex EqualHashRule) []EqualHashRule
}
