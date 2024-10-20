package hash_test

import (
	"testing"

	"github.com/haleyrc/assert"
	"github.com/haleyrc/hash"
)

func TestGenerate(t *testing.T) {
	h := hash.New("hello")
	assert.OK(t, h.Compare("hello"))
	assert.Error(t, h.Compare("goodbye"), "hash mismatch")
}
