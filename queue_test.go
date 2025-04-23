package mutexqueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	assert := assert.New(t)
	q := New()
	el1 := "foo"
	el2 := "bar"
	el3 := "baz"
	assert.Equal(true, q.IsEmpty())
	q.Put(el1)
	assert.Equal(false, q.IsEmpty())
	q.Put(el2)
	q.Put(el3)
	assert.Equal(3, q.GetSize())
	assert.Equal(el1, q.Get())
	assert.Equal(el2, q.Get())
	assert.Equal(el3, q.Get())
	assert.Equal(true, q.IsEmpty())
}
