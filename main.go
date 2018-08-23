package main

import (
	"./player"
	"./player/manager"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func initPlayerMap() manager.PlayerMap {
	playerMap := manager.MakePlayerMap()

	fp, err := os.Open("./DEF_PLAYER")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		text := scanner.Text()

		if !strings.HasPrefix(text, "#") {
			ok := playerMap.Add(text)
			if !ok {
				panic(fmt.Sprintf("Error: \"%s\" is invalid player definition.\n", text))
			}
		}
	}

	return playerMap
}

func main() {
	playerMap := initPlayerMap()

	var playerNames [2]string
	var playTimes int

	flag.StringVar(&playerNames[0], "player1", "", "First player name")
	flag.StringVar(&playerNames[1], "player2", "", "Second player name")
	flag.IntVar(&playTimes, "t", 1, "Times to play")
	flag.Parse()

	var ok bool
	var players [2]player.Player

	for i := 0; i < len(playerNames); i++ {
		players[i], ok = playerMap.Get(playerNames[i])
		if !ok {
			panic(fmt.Sprintf("Error: The player \"%s\" is not exist.\n", playerNames[i]))
		}
	}

	prints := false
	if playerNames[0] == "human" || playerNames[1] == "human" {
		prints = true
	}

	p1Win, p2Win, draw := PlaySomeGames(players[0], players[1], playTimes, prints)

	fmt.Printf("Play result: P1:P2:Draw = %d:%d:%d\n", p1Win, p2Win, draw)
}
