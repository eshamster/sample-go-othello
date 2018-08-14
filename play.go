package main

import (
	"./game"
	"./move"
	"./player"
	"fmt"
	"math/rand"
	"os"
	"time"
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

// Returns winning count of (player1, player2, draw).
// It can be used for easy checking of player's strength.
func PlaySomeGames(player1, player2 player.Player, playTimes int) (int, int, int) {
	// Temporal testing to check strength of player
	rand.Seed(time.Now().UnixNano())

	players := []player.Player{player1, player2}

	winCounts := [3]int{0, 0, 0} // player1, player2, draw

	for i := 0; i < playTimes; i++ {
		wIndex := 0
		bIndex := 1
		if i%2 == 1 {
			wIndex = 1
			bIndex = 0
		}

		play := MakePlay(players[wIndex], players[bIndex])
		play.MoveToEnd(false)

		white, black := play.game.GetPieceCounts()
		if white > black {
			winCounts[wIndex]++
		} else if white < black {
			winCounts[bIndex]++
		} else {
			winCounts[2]++
		}
	}

	return winCounts[0], winCounts[1], winCounts[2]
}
