package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeStore(t *testing.T) {
	store := NewFakeStore()

	// Test Set and Get methods
	store.Set("key1", Value("value1"))
	store.Set("key2", Value("value2"))

	val1, err1 := store.Get("key1")
	assert.NoError(t, err1)
	assert.Equal(t, Value("value1"), val1)

	val2, err2 := store.Get("key2")
	assert.NoError(t, err2)
	assert.Equal(t, Value("value2"), val2)

	// Test Get method with non-existent key
	_, err := store.Get("non-existent")
	assert.Error(t, err)
}
