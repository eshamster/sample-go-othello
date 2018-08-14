package player

import (
	"../game"
	"../move"
)

// TODO: Enable to limit simulation according to time.

// McPlayer is a player using simple Monte-Carlo method
type McPlayer struct {
	randomPlayer  *RandomPlayer
	simulateTimes int
}

type simulateNode struct {
	winCount, selectCount int
}

type simulateInfo struct {
	moves []move.Move
	nodes []simulateNode
}

// MakeMcPlayer makes a player using simple Monte-Carlo method
func MakeMcPlayer(randomPlayer *RandomPlayer, simulateTimes int) McPlayer {
	return McPlayer{
		randomPlayer:  randomPlayer,
		simulateTimes: simulateTimes,
	}
}

// GetMove selects a move using simple Monte-Carlo method
func (player *McPlayer) GetMove(game *game.Game) move.Move {
	preTurnCount := game.GetTurnCount()
	defer game.GoBackTo(preTurnCount)

	isWhiteTurn := game.IsWhiteTurn()
	legalMoves := game.GetLegalMoves()

	info := simulateInfo{
		moves: legalMoves,
		nodes: make([]simulateNode, len(legalMoves)),
	}

	for i := 0; i < player.simulateTimes; i++ {
		selectIndex := selectNode(game, info.nodes)

		game.MoveGame(info.moves[selectIndex])
		info.nodes[selectIndex].selectCount++
		info.nodes[selectIndex].winCount +=
			simulateOnce(game, player.randomPlayer, isWhiteTurn)
		game.GoBackTo(preTurnCount)
	}

	// TODO: Using average rate is maybe better than UCB value.
	return info.moves[selectNode(game, info.nodes)]
}

func selectNode(game *game.Game, nodes []simulateNode) int {
	simulateCount := 0

	for _, node := range nodes {
		simulateCount += node.selectCount
	}

	maxUcb := float64(-99999)
	maxIndex := 0

	for i, node := range nodes {
		ucb := calcUcbDefault(node.winCount, node.selectCount, simulateCount)

		if ucb > maxUcb {
			maxUcb = ucb
			maxIndex = i
		}
	}

	return maxIndex
}

// Return 1 if win, 0 if draw, -1 if lose.
func simulateOnce(game *game.Game, player *RandomPlayer, isWhiteTurn bool) int {
	preTurnCount := game.GetTurnCount()
	defer game.GoBackTo(preTurnCount)

	for !game.IsGameEnd() {
		// XXX: Should check if MoveGame succeeded or not.
		game.MoveGame(player.GetMove(game))
	}

	white, black := game.GetPieceCounts()
	whiteWinPoint := 1
	if !isWhiteTurn {
		whiteWinPoint = -1
	}

	if white > black {
		return whiteWinPoint
	} else if white < black {
		return whiteWinPoint * -1
	}

	return 0
}
