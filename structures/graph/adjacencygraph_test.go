package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdjacencyGraph_ToArray(t *testing.T) {
	ag := prepareAdjacencyGraph()

	agArr := ag.ToArray()

	assert.NotNil(t, agArr)
	assert.Equal(t, ag.Size(), len(agArr))
}
