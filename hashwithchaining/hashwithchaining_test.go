package hashwithchaining

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hashtable := NewHashTable()
	hashtable.Set("veve", 200)

	assert.Equal(t, hashtable.Get("veve"), 200)
}
