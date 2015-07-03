package kademliago

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	data := []byte("this is the test")
	assert.Equal(t, Hash(data), Hash(data), "Hash function should provide the same result for the same input")
}
