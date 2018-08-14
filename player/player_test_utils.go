package player

import (
	"../game"
	"testing"
)

func testPlayer(t *testing.T, player Player) {
	game := game.MakeGame()

	expectedTurnCount := uint(3)

	for i := uint(0); i < expectedTurnCount; i++ {
		move := player.GetMove(&game)
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

