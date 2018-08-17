package player

import (
	"testing"
)

func TestMinimaxPlayer(t *testing.T) {
	player := MakeMinimaxPlayer(4)
	testPlayer(t, &player)
}
