package board

import (
	"../move"
	"fmt"
	"io"
	"log"
)

// Board is a structure having board information
type Board struct {
	black, white uint64
}

const (
	// Empty square state
	Empty = iota
	// White square state
	White
	// Black square state
	Black
)

// Direction
const (
	top = iota
	topRight
	right
	bottomRight
	bottom
	bottomLeft
	left
	topLeft
)

func (board *Board) getSquareState(move move.Move) int {
	moveBit := move.GetBit()

	if board.black&moveBit != 0 {
		return Black
	} else if board.white&moveBit != 0 {
		return White
	} else {
		return Empty
	}
}

func (board *Board) setMoveRaw(move move.Move, isWhite bool) {
	moveBit := move.GetBit()

	if isWhite {
		board.white |= moveBit
		board.black &= ^moveBit
	} else {
		board.black |= moveBit
		board.white &= ^moveBit
	}
}

// MakeBoard makes a initialized board
func MakeBoard() Board {
	board := Board{}

	board.setMoveRaw(move.MakeMove(3, 4), true)
	board.setMoveRaw(move.MakeMove(4, 3), true)
	board.setMoveRaw(move.MakeMove(3, 3), false)
	board.setMoveRaw(move.MakeMove(4, 4), false)

	return board
}

// CopyBoard copies src board to dst one
func CopyBoard(src, dst *Board) {
	dst.white = src.white
	dst.black = src.black
}

func getNextMoveBit(moveBit uint64, direction int) (uint64, bool) {
	maskToBottom := uint64(0x00ffffffffffffff)
	maskToTop := uint64(0xffffffffffffff00)
	maskToRight := uint64(0x7f7f7f7f7f7f7f7f)
	maskToLeft := uint64(0xfefefefefefefefe)

	result := uint64(0)

	switch direction {
	case top:
		result = (moveBit << 8) & maskToTop
	case topRight:
		result = (moveBit << 7) & maskToTop & maskToRight
	case right:
		result = (moveBit >> 1) & maskToRight
	case bottomRight:
		result = (moveBit >> 9) & maskToBottom & maskToRight
	case bottom:
		result = (moveBit >> 8) & maskToBottom
	case bottomLeft:
		result = (moveBit >> 7) & maskToBottom & maskToLeft
	case left:
		result = (moveBit << 1) & maskToLeft
	case topLeft:
		result = (moveBit << 9) & maskToTop & maskToLeft
	default:
		log.Fatalf("Unrecognized direction: %d", direction)
	}

	return result, result != 0
}

// Find pieces that can be reversed from move point towards the direction.
// Return 0 if no pieces can be reversed.
func (board *Board) findReversePieces(moveBit uint64, isWhite bool, direction int) uint64 {
	myBoard := board.white
	otherBoard := board.black

	if !isWhite {
		myBoard = board.black
		otherBoard = board.white
	}

	result := uint64(0)
	findsMine := false

	targetBit := moveBit
	getsNext := false

	for {
		targetBit, getsNext = getNextMoveBit(targetBit, direction)
		if !getsNext {
			break
		}

		if myBoard&targetBit != 0 {
			findsMine = true
			break
		} else if otherBoard&targetBit != 0 {
			result |= targetBit
		} else { /* empty case */
			break
		}
	}

	if !findsMine || (result == 0) {
		return 0
	}

	return result
}

// SetMove sets move to board with reversing opponent pieces.
// It returns true if the move is legal, otherwise returns false.
func (board *Board) SetMove(move move.Move, isWhite bool) bool {
	moveBit := move.GetBit()

	if board.white&moveBit != 0 ||
		board.black&moveBit != 0 {
		// Not empty case
		return false
	}

	result := false

	for i := 0; i < 8; i++ {
		reversePieces := board.findReversePieces(moveBit, isWhite, i)

		if reversePieces != 0 {
			if isWhite {
				board.white |= reversePieces
				board.black &= ^reversePieces
			} else {
				board.black |= reversePieces
				board.white &= ^reversePieces
			}
			result = true
		}
	}

	if result {
		board.setMoveRaw(move, isWhite)
	}

	return result
}

// IsLegalMove returns the move is legal or not
func (board *Board) IsLegalMove(move move.Move, isWhite bool) bool {
	if board.getSquareState(move) != Empty {
		return false
	}

	moveBit := move.GetBit()

	for i := 0; i < 8; i++ {
		if board.findReversePieces(moveBit, isWhite, i) != 0 {
			return true
		}
	}

	return false
}

// HasLegalMove returns if there is a legal move.
func (board *Board) HasLegalMove(isWhite bool) bool {
	// TODO: Use more sophisticated way to get candidates
	for y := uint(0); y < 8; y++ {
		for x := uint(0); x < 8; x++ {
			if board.IsLegalMove(move.MakeMove(x, y), isWhite) {
				return true
			}
		}
	}
	return false
}

func printSquare(w io.Writer, state int) {
	switch state {
	case Empty:
		fmt.Fprint(w, "-")
	case White:
		fmt.Fprint(w, "o")
	case Black:
		fmt.Fprint(w, "x")
	}
}

// PrintBoard prints board to specified output
func (board *Board) PrintBoard(w io.Writer, printsNumber bool) {
	if printsNumber {
		fmt.Fprintln(w, "  abcdefgh")
	}

	for y := uint(0); y < 8; y++ {
		if printsNumber {
			fmt.Fprintf(w, "%d ", y+1)
		}
		for x := uint(0); x < 8; x++ {
			printSquare(w, board.getSquareState(move.MakeMove(x, y)))
		}
		fmt.Fprint(w, "\n")
	}
}
