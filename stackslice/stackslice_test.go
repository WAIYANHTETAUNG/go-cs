package stackslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	stackobj := NewStackSlice()
	stackobj.push(1)
	assert.True(t, stackobj.isEmpty())
}

func TestPop(t *testing.T) {
	stackobj := NewStackSlice()
	stackobj.push(1)
	stackobj.push(2)
	lastItem := stackobj.pop()

	assert.Equal(t, lastItem, 2)
}
