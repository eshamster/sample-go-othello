package main

import (
	"./player"
	"fmt"
)

func main() {
	randPlayer := player.MakeRandomPlayer(player.DefaultPolicy)
	// player1 := player.MakeRandomPlayer(player.DefaultPolicy)
	player1 := player.MakeMcPlayer(&randPlayer, 100)
	// player2 := player.MakeUctPlayer(&randPlayer, 100)
	// player2 := player.MakeHumanPlayer(player.MakeDefaultHumanMoveSelector())
	player2 := player.MakeMinimaxPlayer(4)

	p1Win, p2Win, draw := PlaySomeGames(&player1, &player2, 10)

	fmt.Printf("P1:P2:D = %d:%d:%d\n", p1Win, p2Win, draw)
}
