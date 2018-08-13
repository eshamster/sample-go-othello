package board

import (
	"../move"
	"bytes"
	"testing"
)

func getBoardString(board *Board, printsNumber bool) string {
	buf := &bytes.Buffer{}
	board.PrintBoard(buf, printsNumber)
	return buf.String()
}

func checkBoardResult(t *testing.T, expectedString string, board *Board) {
	result := getBoardString(board, false)

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

func TestBoardPrintWithNumber(t *testing.T) {
	board := MakeBoard()
	expectedString := `  abcdefgh
1 --------
2 --------
3 --------
4 ---xo---
5 ---ox---
6 --------
7 --------
8 --------
`

	result := getBoardString(&board, true)

	if expectedString != result {
		t.Errorf("Expected:\n%s\nResult:\n%s", expectedString, result)
	}
}

func TestCopyBoard(t *testing.T) {
	board := MakeBoard()
	board.setMoveRaw(move.MakeMove(0, 0), true)

	var dst Board
	CopyBoard(&board, &dst)

	checkBoardResult(t, getBoardString(&board, false), &dst)
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

func TestHasLegalMove(t *testing.T) {
	board := MakeBoard()

	// Test using initial board
	if !board.HasLegalMove(true) {
		t.Errorf("The inital board should have legal moves for white")
	}
	if !board.HasLegalMove(false) {
		t.Errorf("The inital board should have legal moves for black")
	}

	// Modify board
	setPiecesRaw(&board, [][2]uint{
		{2, 2}, {3, 2}, {4, 2}, {5, 2},
		{2, 3}, {5, 3},
		{2, 4}, {5, 4},
		{2, 5}, {3, 5}, {4, 5}, {5, 5},
	}, true)
	setPiecesRaw(&board, [][2]uint{{3, 4}, {4, 3}}, false)

	expected := `--------
--------
--oooo--
--oxxo--
--oxxo--
--oooo--
--------
--------
`
	checkBoardResult(t, expected, &board)

	// Test using modified board
	if board.HasLegalMove(true) {
		t.Errorf("The board should not have legal moves for white")
	}
	if !board.HasLegalMove(false) {
		t.Errorf("The board should have legal moves for black")
	}
}

func checkLegalMoves(t *testing.T, expected, resultMoves []move.Move) {
	checkResult := true

LOOP:
	for _, target := range expected {
		for _, result := range resultMoves {
			if result.Equals(&target) {
				continue LOOP
			}
		}
		checkResult = false
		break LOOP
	}

	if !checkResult || len(expected) != len(resultMoves) {
		t.Errorf("%v should have same moves to %v", resultMoves, expected)
	}
}

func TestGetLegalMoves(t *testing.T) {
	board := MakeBoard()

	checkLegalMoves(t, []move.Move{
		move.MakeMove(2, 3),
		move.MakeMove(3, 2),
		move.MakeMove(4, 5),
		move.MakeMove(5, 4),
	}, board.GetLegalMoves(true))

	checkLegalMoves(t, []move.Move{
		move.MakeMove(2, 4),
		move.MakeMove(3, 5),
		move.MakeMove(4, 2),
		move.MakeMove(5, 3),
	}, board.GetLegalMoves(false))
}

func TestGetPieceCounts(t *testing.T) {
	board := MakeBoard()
	retval := board.SetMove(move.MakeMove(2, 3), true)

	if !retval {
		t.Errorf("Failed to set move to (2, 3)")
	}

	white, black := board.GetPieceCounts()
	expectedWhite := 4
	expectedBlack := 1

	if white != expectedWhite || black != expectedBlack {
		t.Errorf("(white, black): (%d, %d) != (%d, %d)\n",
			expectedWhite, expectedBlack,
			white, black)
	}
}
