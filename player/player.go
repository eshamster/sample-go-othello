package player

import (
	"../game"
	"../move"
)

// Player moves game one by one.
type Player interface {
	// GetMove returns move
	GetMove(game *game.Game) move.Move
}
