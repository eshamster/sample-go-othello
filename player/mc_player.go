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

type nodeInfo struct {
	winCount, selectCount int
}

type simulateInfo struct {
	moves     []move.Move
	nodeInfos []nodeInfo
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
		moves:     legalMoves,
		nodeInfos: make([]nodeInfo, len(legalMoves)),
	}

	for i := 0; i < player.simulateTimes; i++ {
		selectIndex := player.selectNode(game, &info)

		game.MoveGame(info.moves[selectIndex])
		info.nodeInfos[selectIndex].selectCount++
		info.nodeInfos[selectIndex].winCount += player.simulateOnce(game, isWhiteTurn)
		game.GoBackTo(preTurnCount)
	}

	// TODO: Using average rate is maybe better than UCB value.
	return info.moves[player.selectNode(game, &info)]
}

func (player *McPlayer) selectNode(game *game.Game, info *simulateInfo) int {
	simulateCount := 0

	for _, node := range info.nodeInfos {
		simulateCount += node.selectCount
	}

	maxUcb := float64(-99999)
	maxIndex := 0

	for i, node := range info.nodeInfos {
		ucb := calcUcbDefault(node.winCount, node.selectCount, simulateCount)

		if ucb > maxUcb {
			maxUcb = ucb
			maxIndex = i
		}
	}

	return maxIndex
}

// Return 1 if win, 0 if draw, -1 if lose.
func (player *McPlayer) simulateOnce(game *game.Game, isWhiteTurn bool) int {
	preTurnCount := game.GetTurnCount()
	defer game.GoBackTo(preTurnCount)

	for !game.IsGameEnd() {
		// XXX: Should check if MoveGame succeeded or not.
		game.MoveGame(player.randomPlayer.GetMove(game))
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
