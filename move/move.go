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
	return Move{move: 1 << (x + y*8)}
}

// GetX gets x position of the move
func (move *Move) GetX() uint {
	return uint(bits.TrailingZeros64(move.move)) % 8
}

// GetY gets y position of the move
func (move *Move) GetY() uint {
	return uint(bits.TrailingZeros64(move.move)) / 8
}

// GetBit gets a move by bit flag style
func (move *Move) GetBit() uint64 {
	return move.move
}

// Equals checks if the both moves are same
func (move *Move) Equals(target *Move) bool {
	return move.move == target.move
}
