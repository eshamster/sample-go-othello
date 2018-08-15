package player

import (
	"../game"
	"../move"
	"math/rand"
	"testing"
)

func SelectHumanMoveRandomly(humanMoves []string, _ *game.Game) string {
	return humanMoves[rand.Int()%len(humanMoves)]
}

func TestHumanPlayer(t *testing.T) {
	player := MakeHumanPlayer(SelectHumanMoveRandomly)
	testPlayer(t, &player)
}

func TestConvertMoveToHumanMove(t *testing.T) {
	expected := "c7"
	actual := convertMoveToHumanMove(move.MakeMove(2, 6))

	if expected != actual {
		t.Errorf("%s != %s\n", expected, actual)
	}
}

func TestConvertHumanMoveToMove(t *testing.T) {
	// normal cases
	{
		expected := move.MakeMove(2, 6)
		actual, ok := convertHumanMoveToMove("c7")

		if !ok || !expected.Equals(&actual) {
			t.Errorf("ok: %t, %v != %v", ok, expected, actual)
		}
	}
	{
		expected := move.MakeMove(2, 6)
		actual, ok := convertHumanMoveToMove("C7")

		if !ok || !expected.Equals(&actual) {
			t.Errorf("ok: %t, %v != %v", ok, expected, actual)
		}
	}

	// error cases
	for _, str := range []string{"7c", "c77", "x7", "c9"} {
		move, ok := convertHumanMoveToMove(str)

		if ok {
			t.Errorf("Converting \"%s\" should fail. (return value: %v)", str, move)
		}
	}
}
