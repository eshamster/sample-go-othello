package player

import (
	"../game"
	"testing"
)

func TestRandomPlayer(t *testing.T) {
	game := game.MakeGame()
	whitePlayer := MakeRandomPlayer(DefaultPolicy)

	expectedTurnCount := uint(3)

	for i := uint(0); i < expectedTurnCount; i++ {
		move := whitePlayer.GetMove(&game)
		moveResult := game.MoveGame(move)

		if !moveResult {
			t.Errorf("Illegal move %d, %d", move.GetX(), move.GetY())
		}
	}

	actualTurnCount := game.GetTurnCount()
	if expectedTurnCount != actualTurnCount {
		t.Errorf("Turn count is not right: %d != %d", expectedTurnCount, actualTurnCount)
	}
}
