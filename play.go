package main

import (
	"./game"
	"./move"
	"./player"
	"fmt"
	"os"
)

// Play manages a game play.
type Play struct {
	whitePlayer, blackPlayer player.Player
	game                     game.Game
}

// MakePlay makes a initial play state.
func MakePlay(whitePlayer, blackPlayer player.Player) Play {
	return Play{
		whitePlayer: whitePlayer,
		blackPlayer: blackPlayer,
		game:        game.MakeGame(),
	}
}

// MoveOnce goes to next turn by move.
func (play *Play) MoveOnce() bool {
	if play.game.IsGameEnd() {
		return false
	}

	preTurnCount := play.game.GetTurnCount()
	var move move.Move

	if play.game.IsWhiteTurn() {
		move = play.whitePlayer.GetMove(&play.game)
	} else {
		move = play.blackPlayer.GetMove(&play.game)
	}

	turnCount := play.game.GetTurnCount()
	if preTurnCount != turnCount {
		fmt.Fprintf(os.Stderr,
			"Error: The turn is changed while selecting a move; before: %d, after: %d\n",
			preTurnCount, turnCount)
		return false
	}

	setResult := play.game.MoveGame(move)
	if !setResult {
		fmt.Fprintf(os.Stderr,
			"Error: The move is illegal; x: %d, y: %d\n",
			move.GetX(), move.GetY())
	}

	return setResult
}

// MoveToEnd moves play to end (or error state).
func (play *Play) MoveToEnd(prints bool) {
	if prints {
		play.game.PrintGame(os.Stdout)
	}

	for play.MoveOnce() {
		if prints {
			play.game.PrintGame(os.Stdout)
		}
	}
}
