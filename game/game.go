package game

import (
	"../board"
	"../move"
	"fmt"
	"io"
	"log"
)

// Turn definition
const (
	white = iota
	black
	gameend
)

type gameSnapshot struct {
	board    board.Board
	turn     int
	moveFrom move.Move
}

// Game has game state and history
type Game struct {
	history   [64]gameSnapshot
	turnCount uint
}

// MakeGame makes a initialized game.
func MakeGame() Game {
	game := Game{}
	snapshot := gameSnapshot{
		board: board.MakeBoard(),
		turn:  white,
	}

	game.history[0] = snapshot
	game.turnCount = 0

	return game
}

// IsWhiteTurn returns if the current turn is white.
func (game *Game) IsWhiteTurn() bool {
	return game.history[game.turnCount].turn == white
}

// GetTurnCount returns turn count that starts from 0.
func (game *Game) GetTurnCount() uint {
	return game.turnCount
}

// IsGameEnd returns if the game has been finished.
func (game *Game) IsGameEnd() bool {
	return game.history[game.turnCount].turn == gameend
}

// MoveGame moves a game one turn.
// Returns true if the move is legal.
func (game *Game) MoveGame(move move.Move) bool {
	if game.IsGameEnd() {
		return false
	}

	// Clone board for next turn
	snapshot := &game.history[game.turnCount]
	nextSnapshot := &game.history[game.turnCount+1]
	board.CopyBoard(
		&snapshot.board,
		&nextSnapshot.board)

	// Set move
	success := nextSnapshot.board.SetMove(move, snapshot.turn == white)
	if !success {
		return false
	}

	// Update nextSnapshot
	nextSnapshot.moveFrom = move

	boolToTurn := func(isWhite bool) int {
		if isWhite {
			return white
		}
		return black
	}

	isWhite := game.IsWhiteTurn()
	if nextSnapshot.board.HasLegalMove(!isWhite) {
		nextSnapshot.turn = boolToTurn(!isWhite)
	} else if nextSnapshot.board.HasLegalMove(isWhite) {
		nextSnapshot.turn = boolToTurn(isWhite)
	} else {
		nextSnapshot.turn = gameend
	}

	// Update turn count
	game.turnCount++

	return true
}

// GoBackTo makes game back to the specified turn.
// Returns true if succeeded.
func (game *Game) GoBackTo(targetTurnCount uint) bool {
	if game.GetTurnCount() <= targetTurnCount {
		return false
	}

	game.turnCount = targetTurnCount
	return true
}

func turnToString(turn int) string {
	switch turn {
	case white:
		return "White (o)"
	case black:
		return "Black (x)"
	case gameend:
		return "Finished"
	default:
		log.Fatalf("Unrecognized Turn: %d", turn)
		return ""
	}
}

// PrintGame prints current game state to specified output
func (game *Game) PrintGame(w io.Writer) {
	snapshot := &game.history[game.GetTurnCount()]
	fmt.Fprintf(w, "Turn: %s, Count: %d\n",
		turnToString(snapshot.turn), game.GetTurnCount())
	snapshot.board.PrintBoard(w, true)
}
