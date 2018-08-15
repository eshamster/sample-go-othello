package player

import (
	"testing"
)

func TestUctPlayer(t *testing.T) {
	randPlayer := MakeRandomPlayer(DefaultPolicy)
	player := MakeUctPlayer(&randPlayer, 100)
	testPlayer(t, &player)
}
