package manager

import (
	"../../player"
	"strconv"
	"strings"
)

// PlayerMap is a type to store player and its name.
type PlayerMap map[string]player.Player

// MakePlayerMap makes a player map.
func MakePlayerMap() PlayerMap {
	return PlayerMap{}
}

// Add adds player to player map.
// Syntax:
//   <name> human
//   <name> random
//   <name> minimax <search depth>
//   <name> mc <simulation times>
//   <name> uct <simulation times>
func (playerMap PlayerMap) Add(definition string) bool {
	name, result, ok := parseOneLine(definition)

	if ok {
		playerMap[name] = result
	}

	return ok
}

// Get gets player from player map.
func (playerMap PlayerMap) Get(name string) (player.Player, bool) {
	result, ok := playerMap[name]
	return result, ok
}

func parseOneLine(target string) (string, player.Player, bool) {
	elements := strings.Split(target, " ")

	if len(elements) < 2 {
		return "", nil, false
	}

	name := elements[0]
	kind := elements[1]

	switch kind {
	case "human":
		player := player.MakeHumanPlayer(player.MakeDefaultHumanMoveSelector())
		return name, &player, true
	case "mc":
		if len(elements) < 3 {
			return "", nil, false
		}
		simulateNum, ok := strconv.Atoi(elements[2])
		if ok != nil {
			return "", nil, false
		}
		randomPlayer := player.MakeRandomPlayer(player.DefaultPolicy)
		player := player.MakeMcPlayer(&randomPlayer, simulateNum)
		return name, &player, true
	case "minimax":
		if len(elements) < 3 {
			return "", nil, false
		}
		searchDepth, ok := strconv.Atoi(elements[2])
		if ok != nil {
			return "", nil, false
		}
		player := player.MakeMinimaxPlayer(searchDepth)
		return name, &player, true
	case "random":
		player := player.MakeRandomPlayer(player.DefaultPolicy)
		return name, &player, true
	case "uct":
		if len(elements) < 3 {
			return "", nil, false
		}
		simulateNum, ok := strconv.Atoi(elements[2])
		if ok != nil {
			return "", nil, false
		}
		randomPlayer := player.MakeRandomPlayer(player.DefaultPolicy)
		player := player.MakeUctPlayer(&randomPlayer, simulateNum)
		return name, &player, true
	}

	return "", nil, false
}
