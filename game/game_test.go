package game

import (
	"../board"
	"../move"
	"bytes"
	"testing"
)

// TODO: Test GameEnd case

func getGameString(game *Game) string {
	buf := &bytes.Buffer{}
	game.PrintGame(buf)
	return buf.String()
}

func checkGameState(t *testing.T, expectedString string, game *Game) {
	result := getGameString(game)

	if expectedString != result {
		t.Errorf("Expected:\n%s\nResult:\n%s", expectedString, result)
	}
}

func TestPrintGame(t *testing.T) {
	game := MakeGame()
	expected := `Turn: White (o), Count: 0
  abcdefgh
1 --------
2 --------
3 --------
4 ---xo---
5 ---ox---
6 --------
7 --------
8 --------
`

	checkGameState(t, expected, &game)
}

func TestSetAndGoBack(t *testing.T) {
	game := MakeGame()

	game.MoveGame(move.MakeMove(2, 3)) // c4
	game.MoveGame(move.MakeMove(4, 2)) // e3
	game.MoveGame(move.MakeMove(5, 4)) // f5
	game.MoveGame(move.MakeMove(1, 3)) // b4
	game.MoveGame(move.MakeMove(3, 2)) // d3

	expected := `Turn: Black (x), Count: 5
  abcdefgh
1 --------
2 --------
3 ---ox---
4 -xxoo---
5 ---ooo--
6 --------
7 --------
8 --------
`

	checkGameState(t, expected, &game)

	// Illegal move case
	moved := game.MoveGame(move.MakeMove(0, 0))
	if moved {
		t.Errorf("Move(0, 0) should be illegal move in this state")
	}

	// GoBackTo
	game.GoBackTo(uint(2))
	expected = `Turn: White (o), Count: 2
  abcdefgh
1 --------
2 --------
3 ----x---
4 --oox---
5 ---ox---
6 --------
7 --------
8 --------
`

	checkGameState(t, expected, &game)
}

func TestGetPieceCounts(t *testing.T) {
	game := MakeGame()
	retval := game.MoveGame(move.MakeMove(2, 3))

	if !retval {
		t.Errorf("Failed to set move to (2, 3)")
	}

	white, black := game.GetPieceCounts()
	expectedWhite := 4
	expectedBlack := 1

	if white != expectedWhite || black != expectedBlack {
		t.Errorf("(white, black): (%d, %d) != (%d, %d)\n",
			expectedWhite, expectedBlack,
			white, black)
	}
}

func checkSquareState(t *testing.T, game *Game, expected int, move move.Move) {
	actual := game.GetSquareState(move)

	if expected != actual {
		t.Errorf("Move(%d, %d) should be %s (=%d): actual: %s (=%d)",
			move.GetX(), move.GetY(),
			board.ConvertSquareStateToString(expected), expected,
			board.ConvertSquareStateToString(actual), actual)
	}
}

func TestGetSquareState(t *testing.T) {
	game := MakeGame()

	checkSquareState(t, &game, board.White, move.MakeMove(3, 4))
	checkSquareState(t, &game, board.Black, move.MakeMove(4, 4))
	checkSquareState(t, &game, board.Empty, move.MakeMove(2, 2))
}
