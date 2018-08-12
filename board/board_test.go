package board

import (
	"../move"
	"bytes"
	"testing"
)

func getBoardString(board *Board) string {
	buf := &bytes.Buffer{}
	board.PrintBoard(buf)
	return buf.String()
}

func checkBoardResult(t *testing.T, expectedString string, board *Board) {
	result := getBoardString(board)

	if expectedString != result {
		t.Errorf("Expected:\n%s\nResult:\n%s", expectedString, result)
	}
}

func TestBoardPrint(t *testing.T) {
	board := MakeBoard()
	expected := `--------
--------
--------
---xo---
---ox---
--------
--------
--------
`
	checkBoardResult(t, expected, &board)
}

func TestCopyBoard(t *testing.T) {
	board := MakeBoard()
	board.setMoveRaw(move.MakeMove(0, 0), true)

	var dst Board
	CopyBoard(&board, &dst)

	checkBoardResult(t, getBoardString(&board), &dst)
}

func setPiecesRaw(board *Board, pieces [][2]uint, isWhite bool) {
	for _, piece := range pieces {
		x := piece[0]
		y := piece[1]

		board.setMoveRaw(move.MakeMove(x, y), isWhite)
	}
}

func TestSetMove(t *testing.T) {
	var board Board
	var expected string
	var retval bool

	// Reverse multiple piece and multiple direction
	board = MakeBoard()
	setPiecesRaw(&board, [][2]uint{{4, 6}}, true)
	setPiecesRaw(&board, [][2]uint{{4, 4}, {5, 4}, {4, 5}, {5, 5}, {3, 7}}, false)
	expected = `--------
--------
--------
---xo---
---oxx--
----xx--
----o---
---x----
`
	checkBoardResult(t, expected, &board)

	retval = board.SetMove(move.MakeMove(6, 4), true)

	if !retval {
		t.Errorf("Failed to set move to (6,4)")
	}

	expected = `--------
--------
--------
---xo---
---oooo-
----xo--
----o---
---x----
`
	checkBoardResult(t, expected, &board)

	// Stop by empty or wall
	board = MakeBoard()
	setPiecesRaw(&board, [][2]uint{{4, 6}}, true)
	setPiecesRaw(&board, [][2]uint{{3, 4}, {5, 4}, {7, 4}, {4, 5}, {5, 5}}, false)
	expected = `--------
--------
--------
---xo---
---xxx-x
----xx--
----o---
--------
`
	checkBoardResult(t, expected, &board)

	retval = board.SetMove(move.MakeMove(6, 4), true)

	if !retval {
		t.Errorf("Failed to set move to (6,4)")
	}

	expected = `--------
--------
--------
---xo---
---xxxox
----xo--
----o---
--------
`
	checkBoardResult(t, expected, &board)
}

func testIsLegal(t *testing.T, x, y uint, isWhite, expected bool) {
	board := MakeBoard()
	result := board.IsLegalMove(move.MakeMove(x, y), isWhite)

	if expected != result {
		if expected {
			t.Errorf("Move(%d,%d) should be legal", x, y)
		} else {
			t.Errorf("Move(%d,%d) should be illegal", x, y)
		}
	}
}

func TestIsLegalMove(t *testing.T) {
	// legal case
	testIsLegal(t, 2, 3, true, true)
	testIsLegal(t, 3, 5, false, true)
	// illegal: next to my piece
	testIsLegal(t, 3, 5, true, false)
	testIsLegal(t, 2, 3, false, false)
	// illegal: isolated case
	testIsLegal(t, 1, 1, true, false)
	testIsLegal(t, 1, 1, false, false)
	// illegal: on exsiting piece
	testIsLegal(t, 3, 3, true, false)
	testIsLegal(t, 3, 3, false, false)
	// TODO: illegal: out of board
}
