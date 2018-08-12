package game

import (
	"../move"
	"bytes"
	"testing"
)

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

// TODO: Test GameEnd case
