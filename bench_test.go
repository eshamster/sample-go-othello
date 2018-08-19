package main

import (
	"./player"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	player1 := player.MakeMinimaxPlayer(4)
	randPlayer := player.MakeRandomPlayer(player.DefaultPolicy)
	player2 := player.MakeUctPlayer(&randPlayer, 10000)

	for i := 0; i < b.N; i++ {
		play := MakePlay(&player1, &player2)
		play.MoveOnce()
		play.MoveOnce()
	}
}
