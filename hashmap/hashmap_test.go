package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hashtable := NewHashTable()
	assert.Equal(t, hashtable.Get("veve"), 200)
}
