package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayStack_Push(t *testing.T) {
	st := NewArrayStack()

	st.Push(cInt(32))
	st.Push(cInt(27))
	st.Push(cInt(55))

	assert.Equal(t, 3, st.Size())
}

func TestArrayStack_Pop(t *testing.T) {
	st := prepareArrayStack()

	f1, isF1 := st.Pop()
	f2, isF2 := st.Pop()
	f3, isF3 := st.Pop()

	assert.True(t, isF1)
	assert.True(t, isF2)
	assert.True(t, isF3)
	assert.NotNil(t, f1)
	assert.NotNil(t, f2)
	assert.NotNil(t, f3)
	assert.True(t, f1.Equal(cInt(20)))
	assert.True(t, f2.Equal(cInt(19)))
	assert.True(t, f3.Equal(cInt(18)))
}

func TestArrayStack_Size(t *testing.T) {
	st := prepareArrayStack()

	st.Push(cInt(555))
	st.Push(cInt(444))
	st.Pop()

	assert.Equal(t, prepareArrayStack().Size()+1, st.Size())
}
