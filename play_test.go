package main

import (
	"./game"
	"./player"
	"testing"
)

func checkGameState(t *testing.T, game *game.Game, expectedIsWhite bool, expectedTurnCount uint) {
	actualIsWhite := game.IsWhiteTurn()
	actualTurnCount := game.GetTurnCount()

	if (expectedIsWhite != actualIsWhite) ||
		(expectedTurnCount != actualTurnCount) {
		t.Errorf("Expected game state: IsWhite=%t, TurnCount=%d\nActual: IW=%t, TC=%d",
			expectedIsWhite, expectedTurnCount, actualIsWhite, actualTurnCount)
	}
}

func TestPlay(t *testing.T) {
	whitePlayer := player.MakeRandomPlayer(player.DefaultPolicy)
	blackPlayer := player.MakeRandomPlayer(player.DefaultPolicy)

	play := MakePlay(&whitePlayer, &blackPlayer)
	game := &play.game

	checkGameState(t, game, true, 0)

	play.MoveOnce()
	checkGameState(t, game, false, 1)
	play.MoveOnce()
	checkGameState(t, game, true, 2)

	play.MoveToEnd(false)
	if !game.IsGameEnd() {
		t.Errorf("Game should be finished after play.MoveToEnd")
	}
}
