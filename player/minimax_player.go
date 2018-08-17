package player

import (
	"../board"
	"../game"
	"../move"
)

const (
	maxEvaluateValue = 99999
)

// MinimaxPlayer selects move according to minimax strategy.
type MinimaxPlayer struct {
	searchDepth int
}

// MakeMinimaxPlayer makes a player that selects move according to minimax strategy.
func MakeMinimaxPlayer(searchDepth int) MinimaxPlayer {
	return MinimaxPlayer{
		searchDepth: searchDepth,
	}
}

type minimaxNode struct {
	children  []minimaxNode
	moves     []move.Move
	parent    *minimaxNode
	isWhite   bool
	evalValue int
}

// GetMove selects a move according to minimax strategy.
func (player *MinimaxPlayer) GetMove(game *game.Game) move.Move {
	preTurnCount := game.GetTurnCount()
	defer game.GoBackTo(preTurnCount)

	rootNode := new(minimaxNode)
	player.evaluateRecursively(game, rootNode, 0)

	return rootNode.moves[rootNode.getMaxChildIndex()]
}

func (player *MinimaxPlayer) evaluateRecursively(game *game.Game, node *minimaxNode, depth int) {
	if game.IsGameEnd() || depth >= player.searchDepth {
		node.evalValue = evaluateBoard(game)
		// TODO: Error if node.parent is nil.
		// Such case can occur when GetMove is called in game.IsGameEnd() == true.
		if !node.parent.isWhite {
			node.evalValue *= -1
		}
		return
	}

	preTurnCount := game.GetTurnCount()

	node.isWhite = game.IsWhiteTurn()
	node.moves = game.GetLegalMoves()
	node.children = make([]minimaxNode, len(node.moves))

	for i, move := range node.moves {
		game.MoveGame(move)
		node.children[i].parent = node
		player.evaluateRecursively(game, &node.children[i], depth+1)
		game.GoBackTo(preTurnCount)
	}

	node.evalValue = node.children[node.getMaxChildIndex()].evalValue
	if node.parent != nil &&
		node.isWhite != node.parent.isWhite {
		node.evalValue *= -1
	}
}

func (node *minimaxNode) getMaxChildIndex() int {
	maxEval := -maxEvaluateValue
	result := 0

	for i, child := range node.children {
		if child.evalValue > maxEval {
			maxEval = child.evalValue
			result = i
		}
	}

	return result
}

var cornerMoves []move.Move = []move.Move{
	move.MakeMove(0, 0),
	move.MakeMove(0, 7),
	move.MakeMove(7, 0),
	move.MakeMove(7, 7),
}

const (
	cornerValue = 4
)

// evaluateBoard returns a evaluation value from the view of white player.
func evaluateBoard(game *game.Game) int {
	if game.IsGameEnd() {
		white, black := game.GetPieceCounts()
		if white > black {
			return maxEvaluateValue
		} else {
			return -maxEvaluateValue
		}
	}

	legalMoveCount := len(game.GetLegalMoves())
	if !game.IsWhiteTurn() {
		legalMoveCount *= -1
	}

	sumCornerValue := 0
	for _, move := range cornerMoves {
		state := game.GetSquareState(move)

		if state == board.White {
			sumCornerValue += cornerValue
		} else if state == board.Black {
			sumCornerValue -= cornerValue
		}
	}

	return legalMoveCount + sumCornerValue
}
