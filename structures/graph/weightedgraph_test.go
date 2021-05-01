package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWeightedGraph_AddVertex(t *testing.T) {
	wg := prepareWeightedGraph()

	wg.AddVertex(cInt(51))
	wg.AddVertex(cInt(31))
	wg.AddVertex(cInt(44))
	wg.AddVertex(cInt(44))
	wg.AddVertex(cInt(51))

	assert.Equal(t, prepareWeightedGraph().Size()+3, wg.Size())
}

func TestWeightedGraph_AddEdge(t *testing.T) {
	wg := prepareWeightedGraph()

	e1 := wg.AddEdge(cInt(13), cInt(15), 5)
	e2 := wg.AddEdge(cInt(11), cInt(3), 7)
	e3 := wg.AddEdge(cInt(13), cInt(41), 14)

	assert.True(t, e1)
	assert.True(t, e2)
	assert.False(t, e3)
}

func TestWeightedGraph_Contains(t *testing.T) {
	wg := prepareWeightedGraph()

	c1 := wg.Contains(cInt(31))
	c2 := wg.Contains(cInt(14))
	c3 := wg.Contains(cInt(3))

	assert.False(t, c1)
	assert.True(t, c2)
	assert.True(t, c3)
}

func TestWeightedGraph_GetEdges(t *testing.T) {
	wg := prepareWeightedGraph()

	e1 := wg.GetEdges(cInt(14))
	e2 := wg.GetEdges(cInt(20))
	e3 := wg.GetEdges(cInt(45))

	assert.NotNil(t, e1)
	assert.NotNil(t, e2)
	assert.Nil(t, e3)
}

func TestWeightedGraph_RemoveVertex(t *testing.T) {
	wg := prepareWeightedGraph()

	r1 := wg.RemoveVertex(cInt(32))
	r2 := wg.RemoveVertex(cInt(11))
	r3 := wg.RemoveVertex(cInt(4))

	assert.False(t, r1)
	assert.True(t, r3)
	assert.True(t, r2)
	assert.False(t, wg.Contains(cInt(11)))
	assert.False(t, wg.Contains(cInt(4)))
}

func TestWeightedGraph_RemoveEdge(t *testing.T) {
	wg := prepareWeightedGraph()

	r1 := wg.RemoveEdge(cInt(12), cInt(11))
	r2 := wg.RemoveEdge(cInt(4), cInt(3))
	r3 := wg.RemoveEdge(cInt(12), cInt(51))

	assert.True(t, r1)
	assert.True(t, r2)
	assert.False(t, r3)
}

func TestWeightedGraph_ToArray(t *testing.T) {
	wg := prepareWeightedGraph()

	arr := wg.ToArray()

	assert.True(t, arr != nil)
	assert.True(t, len(arr) == 20)
}

func TestWeightedGraph_Size(t *testing.T) {
	wg := prepareWeightedGraph()

	wg.AddVertex(cInt(32))
	wg.AddVertex(cInt(41))
	wg.AddVertex(cInt(4))

	assert.Equal(t, prepareWeightedGraph().Size() + 2, wg.Size())
}