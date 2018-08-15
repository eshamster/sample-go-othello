package player

import (
	"../game"
	"../move"
	"fmt"
)

// UCT = Upper Confidence bound applied to Trees

// UctPlayer is a player using UCT
type UctPlayer struct {
	randomPlayer   *RandomPlayer
	simulateTimes  int
	expandInterval int
}

var nextUcdNodeID int

type uctNode struct {
	id         int // for debug
	sumProfit  int
	visitCount int
	parent     *uctNode
	isWhite    bool
	moves      []move.Move
	children   []uctNode
}

func makeUctNode(game *game.Game, parent *uctNode) uctNode {
	moves := game.GetLegalMoves()

	result := uctNode{
		id:       nextUcdNodeID,
		isWhite:  game.IsWhiteTurn(),
		moves:    moves,
		children: make([]uctNode, 0, len(moves)),
		parent:   parent,
	}
	nextUcdNodeID++

	return result
}

// MakeUctPlayer makes a player using simple Monte-Carlo method
func MakeUctPlayer(randomPlayer *RandomPlayer, simulateTimes int) UctPlayer {
	return UctPlayer{
		randomPlayer:   randomPlayer,
		simulateTimes:  simulateTimes,
		expandInterval: 3,
	}
}

// GetMove selects a move using UCT
func (player *UctPlayer) GetMove(game *game.Game) move.Move {
	preTurnCount := game.GetTurnCount()
	defer game.GoBackTo(preTurnCount)

	rootNode := makeUctNode(game, nil)

	for i := 0; i < player.simulateTimes; i++ {
		player.simulateUctOnce(game, &rootNode)
	}

	// TODO: Using average rate is maybe better than UCB value.
	return rootNode.moves[rootNode.selectUctChild(game)]
}

func (player *UctPlayer) simulateUctOnce(game *game.Game, rootNode *uctNode) {
	preTurnCount := game.GetTurnCount()
	defer game.GoBackTo(preTurnCount)

	leafNode := player.selectToLeaf(game, rootNode)
	whiteProfit := simulateOnce(game, player.randomPlayer, true)
	player.propagateProfitToBackward(game, leafNode, whiteProfit)
}

func (player *UctPlayer) selectToLeaf(game *game.Game, node *uctNode) *uctNode {
	// Node: node.visitCount will be updated in backpropagation
	player.expandUctNodeIfRequired(game, node)

	if node.isLeaf() {
		return node
	}

	index := node.selectUctChild(game)
	game.MoveGame(node.moves[index])
	return player.selectToLeaf(game, &node.children[index])
}

// TODO: Commonize the definition with "selectNode" in mc_player
func (node *uctNode) selectUctChild(game *game.Game) int {
	simulateCount := node.visitCount

	maxUcb := float64(-99999)
	maxIndex := 0

	for i, child := range node.children {
		ucb := calcUcbDefault(child.sumProfit, child.visitCount, simulateCount)

		if ucb > maxUcb {
			maxUcb = ucb
			maxIndex = i
		}
	}

	return maxIndex
}

func (node *uctNode) isRoot() bool {
	return node.parent == nil
}

func (node *uctNode) isLeaf() bool {
	return len(node.children) == 0
}

func (player *UctPlayer) expandUctNodeIfRequired(game *game.Game, node *uctNode) {
	// Note: In the case where the expanded node is selected just after expanding,
	// going back game can be omitted.
	preTurnCount := game.GetTurnCount()
	defer game.GoBackTo(preTurnCount)

	if player.shouldExpandNode(game, node) {
		move := node.moves[len(node.children)]
		game.MoveGame(move)
		node.children = append(node.children, makeUctNode(game, node))
	}
}

func (player *UctPlayer) shouldExpandNode(game *game.Game, node *uctNode) bool {
	return (len(node.children) < len(node.moves)) &&
		((node.isRoot() && node.isLeaf()) ||
			node.visitCount%player.expandInterval == player.expandInterval-1)
}

func (player *UctPlayer) propagateProfitToBackward(game *game.Game, leafNode *uctNode, whiteProfit int) {
	for node := leafNode; !node.isRoot(); node = node.parent {
		profit := whiteProfit
		if !node.parent.isWhite {
			profit *= -1
		}

		node.visitCount++
		node.sumProfit += profit

		if node.parent.isRoot() {
			node.parent.visitCount++
			node.parent.sumProfit += whiteProfit // only for debug
		}
	}
}

func (node *uctNode) print(depth int, maxDepth int) {
	if depth >= maxDepth {
		return
	}

	for i := 0; i < depth; i++ {
		fmt.Printf("  ")
	}

	turn := "White"
	if !node.isWhite {
		turn = "Black"
	}
	fmt.Printf("%s:%d/%d\n", turn, node.sumProfit, node.visitCount)

	for _, child := range node.children {
		child.print(depth+1, maxDepth)
	}
}
