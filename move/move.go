package move

import (
	"math/bits"
)

// Move is a move (x, y position) of Othello
type Move struct {
	move uint64
}

// MakeMove makes a move by x and y positions
func MakeMove(x, y uint) Move {
	return Move{move: 1 << (y + x*8)}
}

// GetMoveX gets x position of the move
func GetMoveX(move Move) uint {
	return uint(bits.TrailingZeros64(move.move)) / 8
}

// GetMoveY gets y position of the move
func GetMoveY(move Move) uint {
	return uint(bits.TrailingZeros64(move.move)) % 8
}

// GetMoveBit gets a move by bit flag style
func GetMoveBit(move Move) uint64 {
	return move.move
}
