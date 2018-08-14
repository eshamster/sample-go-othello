package main

import (
	"./player"
	"fmt"
)

func main() {
	player1 := player.MakeRandomPlayer(player.DefaultPolicy)
	randPlayer := player.MakeRandomPlayer(player.DefaultPolicy)
	player2 := player.MakeMcPlayer(&randPlayer, 100)

	p1Win, p2Win, draw := PlaySomeGames(&player1, &player2, 11)

	fmt.Printf("P1:P2:D = %d:%d:%d\n", p1Win, p2Win, draw)
}
