package main

import (
	"./player/manager"
	"bufio"
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

	fmt.Printf("%#v\n", playerMap)
}
