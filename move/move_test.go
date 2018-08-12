package move

import (
	"testing"
)

func TestGetMovePosition(t *testing.T) {
	move := MakeMove(3, 5)
	expectX := uint(3)
	actualX := move.GetX()
	expectY := uint(5)
	actualY := move.GetY()
	expectBit := uint64(1 << (3 + 5*8))
	actualBit := move.GetBit()

	if expectX != actualX {
		t.Errorf("X: %d != %d", expectX, actualX)
	}
	if expectY != actualY {
		t.Errorf("Y: %d != %d", expectY, actualY)
	}
	if expectBit != actualBit {
		t.Errorf("Bit: %d != %d", expectBit, actualBit)
	}
}
