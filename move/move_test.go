package move

import (
	"testing"
)

func TestGetMovePosition(t *testing.T) {
	move := MakeMove(3, 5)
	expectX := uint(3)
	actualX := GetMoveX(move)
	expectY := uint(5)
	actualY := GetMoveY(move)

	if expectX != actualX {
		t.Errorf("X: %d != %d", expectX, actualX)
	}
	if expectY != actualY {
		t.Errorf("Y: %d != %d", expectY, actualY)
	}
}
