package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventHandler(t *testing.T) {
	handler := NewEventHandler(2)
	called := false
	callback := func(a interface{}) {
		called = true
	}
	handler.On(1, callback)

	handler.Trigger(1, "CENAS")
	assert.True(t, called)
}
