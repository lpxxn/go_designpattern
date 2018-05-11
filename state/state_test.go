package state

import (
	"testing"
)

func TestState(t *testing.T) {
	c := NewContext()
	c.DoAction()
	c.DoAction()
	c.DoAction()
}