package player

import (
	"../game"
	"../move"
	"log"
	"math/rand"
)

// GetPolicy returns policy.
// The policy is probability disribution that represents which legal
// moves will be selected in what probability.
type GetPolicy func(legalMoves []move.Move, game *game.Game) []float64

// RandomPlayer moves randomly
type RandomPlayer struct {
	getPolicy GetPolicy
}

// MakeRandomPlayer makes a random player
func MakeRandomPlayer(getPolicy GetPolicy) RandomPlayer {
	return RandomPlayer{getPolicy: getPolicy}
}

// DefaultPolicy is uniform random policy
func DefaultPolicy(legalMoves []move.Move, _ *game.Game) []float64 {
	// Note: If the allocation of result will be bottleneck,
	// the cost can be reduced by caching (or generating in advance)
	// the result in such map as map[int][]float64.
	// (The "int" is length of the legalMoves, and the "[]float64" is result)
	len := len(legalMoves)
	result := make([]float64, len)

	for i := range legalMoves {
		result[i] = 1.0 / float64(len)
	}

	return result
}

// GetMove ramdonly
func (player *RandomPlayer) GetMove(game *game.Game) move.Move {
	legalMoves := game.GetLegalMoves()
	policy := player.getPolicy(legalMoves, game)

	randVal := rand.Float64() // [0, 1)
	sumProb := float64(0)
	len := len(policy)

	for i, prob := range policy {
		sumProb += prob
		if sumProb > randVal || i == len-1 {
			return legalMoves[i]
		}
	}

	// Should not reach
	log.Fatalf("Should not reach here: rand = %f, policy = %v\n", randVal, policy)
	return move.MakeMove(0, 0)
}
