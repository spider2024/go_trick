package simple

import "testing"

func TestCooke(t *testing.T) {
	xhs := &Xhs{}
	doCook(xhs)

	egg := &Egg{}
	doCook(egg)
}
