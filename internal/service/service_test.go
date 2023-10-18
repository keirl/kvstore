package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_New(t *testing.T) {
	store := NewFakeStore()
	s := NewService(store)

	assert.NotNil(t, s)
}
