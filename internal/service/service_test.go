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

func TestService_SetV(t *testing.T) {
	store := NewFakeStore()
	s := NewService(store)

	cmd := SetVCmd{
		Key:     "key",
		Version: 1,
		Value:   Value("value"),
	}

	err := s.SetV(cmd)
	assert.Nil(t, err)
}
