package player

import (
	"../game"
	"../move"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// SelectHumanMove is function to select human move from candidates.
type SelectHumanMove func([]string, *game.Game) string

// HumanPlayer is player controlled by human.
type HumanPlayer struct {
	selectHumanMove SelectHumanMove
}

// MakeHumanPlayer makes a human player
func MakeHumanPlayer(selectHumanMove SelectHumanMove) HumanPlayer {
	return HumanPlayer{
		selectHumanMove: selectHumanMove,
	}
}

// MakeDefaultHumanMoveSelector return default SelectHumanMove.
func MakeDefaultHumanMoveSelector() SelectHumanMove {
	return func(humanMoves []string, game *game.Game) string {
		fmt.Printf("\nLegal Moves: %s\n", strings.Join(humanMoves, ", "))

		result := ""

		for !containsString(result, humanMoves) {
			result = readHumanInputOnce("Please input move")
		}

		return result
	}
}

func readHumanInputOnce(text string) string {
	fmt.Printf("%s: ", text)
	s := bufio.NewScanner(os.Stdin)
	ok := s.Scan()

	if !ok {
		log.Fatalln("Failed to scan from os.Stdin.")
	}

	return s.Text()
}

// GetMove selects a move according to human player
func (player *HumanPlayer) GetMove(game *game.Game) move.Move {
	moves := game.GetLegalMoves()
	humanMoves := make([]string, 0, len(moves))

	for _, move := range moves {
		humanMoves = append(humanMoves, convertMoveToHumanMove(move))
	}

	selectedHumanMove := ""

	// TODO: Exit if len(moves) == 0

	for !containsString(selectedHumanMove, humanMoves) {
		selectedHumanMove = player.selectHumanMove(humanMoves, game)
	}

	// TODO: Check the second value.
	result, _ := convertHumanMoveToMove(selectedHumanMove)

	return result
}

func containsString(str string, stringList []string) bool {
	for _, strInList := range stringList {
		if str == strInList {
			return true
		}
	}
	return false
}

func convertMoveToHumanMove(move move.Move) string {
	x := rune(int('a') + int(move.GetX()))
	y := rune(int('1') + int(move.GetY()))
	return string(x) + string(y)
}

func convertHumanMoveToMove(humanMove string) (move.Move, bool) {
	dummyMove := move.MakeMove(0, 0)

	if len(humanMove) != 2 {
		return dummyMove, false
	}

	xInt := uint(strings.ToLower(humanMove)[0])
	yInt := uint(humanMove[1])

	if xInt < uint('a') || xInt > uint('h') ||
		yInt < uint('1') || yInt > uint('8') {
		return dummyMove, false
	}

	return move.MakeMove(xInt-uint('a'), yInt-uint('1')), true
}
