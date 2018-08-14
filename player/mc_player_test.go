package player

import (
	"testing"
)

func TestMcPlayer(t *testing.T) {
	randPlayer := MakeRandomPlayer(DefaultPolicy)
	player := MakeMcPlayer(&randPlayer, 100)
	testPlayer(t, &player)
}
