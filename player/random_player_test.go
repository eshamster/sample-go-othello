package player

import (
	"testing"
)

func TestRandomPlayer(t *testing.T) {
	player := MakeRandomPlayer(DefaultPolicy)
	testPlayer(t, &player)
}
