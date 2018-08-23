package manager

import (
	"../../player"
	"testing"
)

func checkNg(t *testing.T, ok bool, reason string) {
	if ok {
		t.Errorf("Should ng because %s\n", reason)
	}
}

func checkOkAndName(t *testing.T, ok bool, expectedName, actualName string) {
	if !ok {
		t.Errorf("Should ok\n")
	}
	if expectedName != actualName {
		t.Errorf("Name: \"%s\" != \"%s\"\n", expectedName, actualName)
	}
}

func TestPlayerMap(t *testing.T) {
	playerMap := MakePlayerMap()
	var ok bool
	var somePlayer player.Player

	ok = playerMap.Add("wrong_definition")
	checkNg(t, ok, "passes a wrong definition")

	ok = playerMap.Add("player1 human")
	if !ok {
		t.Errorf("Should ok\n")
	}
	ok = playerMap.Add("player2 uct 10000")
	if !ok {
		t.Errorf("Should ok\n")
	}

	// TODO: Check type of somePlayer
	somePlayer, ok = playerMap.Get("player2")
	if !ok || somePlayer == nil {
		t.Errorf("Should ok\n")
	}

	somePlayer, ok = playerMap.Get("not_defined")
	checkNg(t, ok, "the player \"not_defined\" is not defined")
}

func TestParseOneLine(t *testing.T) {
	var name string
	var result player.Player
	var ok bool

	/* basic error cases */
	name, result, ok = parseOneLine("one_arg")
	checkNg(t, ok, "the number of arguments is short")

	name, result, ok = parseOneLine("not_defined abcd")
	checkNg(t, ok, "player kind \"abcd\" is not defined")

	/* HumanPlayer */
	name, result, ok = parseOneLine("hp human")
	checkOkAndName(t, ok, "hp", name)

	switch v := result.(type) {
	case (*player.HumanPlayer): // do nothing
	default:
		t.Errorf("Type: HumanPlayer != %T", v)
	}

	/* McPlayer */
	name, result, ok = parseOneLine("short mc")
	checkNg(t, ok, "McPlayer requires 1 extra argument")
	name, result, ok = parseOneLine("wront_type mc aaaa")
	checkNg(t, ok, "McPlayer requires number as 1 extra argument")

	name, result, ok = parseOneLine("mp mc 1000")
	checkOkAndName(t, ok, "mp", name)

	switch v := result.(type) {
	case (*player.McPlayer): // do nothing
	default:
		t.Errorf("Type: McPlayer != %T", v)
	}

	/* MinimaxPlayer */
	name, result, ok = parseOneLine("short minimax")
	checkNg(t, ok, "MinimaxPlayer requires 1 extra argument")
	name, result, ok = parseOneLine("wront_type minimax aaaa")
	checkNg(t, ok, "MinimaxPlayer requires number as 1 extra argument")

	name, result, ok = parseOneLine("mmp minimax 6")
	checkOkAndName(t, ok, "mmp", name)

	switch v := result.(type) {
	case (*player.MinimaxPlayer): // do nothing
	default:
		t.Errorf("Type: MinimaxPlayer != %T", v)
	}

	/* RandomPlayer */
	name, result, ok = parseOneLine("rp random")
	checkOkAndName(t, ok, "rp", name)

	switch v := result.(type) {
	case (*player.RandomPlayer): // do nothing
	default:
		t.Errorf("Type: RandomPlayer != %T", v)
	}

	/* UctPlayer */
	name, result, ok = parseOneLine("short uct")
	checkNg(t, ok, "UctPlayer requires 1 extra argument")
	name, result, ok = parseOneLine("wront_type uct aaaa")
	checkNg(t, ok, "UctPlayer requires number as 1 extra argument")

	name, result, ok = parseOneLine("up uct 2000")
	checkOkAndName(t, ok, "up", name)

	switch v := result.(type) {
	case (*player.UctPlayer): // do nothing
	default:
		t.Errorf("Type: UctPlayer != %T", v)
	}
}
