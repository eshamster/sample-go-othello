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

func TestEquals(t *testing.T) {
	move := MakeMove(1, 2)
	moveSame := MakeMove(1, 2)
	moveDiff := MakeMove(2, 1)

	if !move.Equals(&moveSame) {
		t.Errorf("The 2 moves should be same")
	}
	if move.Equals(&moveDiff) {
		t.Errorf("The 2 moves should be different")
	}
}
