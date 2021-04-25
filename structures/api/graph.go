package api

type Graph interface {
	AddVertex(val string)
	AddEdge(fVertex string, sVertex string)
	RemoveEdge(fVertex string, sVertex string)
	RemoveVertex(fVertex string) bool
}
