package player

import (
	"testing"
)

func TestCalcUcb(t *testing.T) {
	if calcUcbDefault(100, 0, 100) != ucbMax {
		t.Error("Should return max value if num == 0")
	}
	if calcUcbDefault(100, 100, 0) != ucbMax {
		t.Error("Should return max value if totalNum == 0")
	}

	if calcUcbDefault(0, 100, 100) > calcUcbDefault(10, 100, 100) {
		t.Error("If num and totalNum are same, bigger sum should cause begger UCB")
	}
	if calcUcbDefault(10, 20, 100) > calcUcbDefault(10, 10, 100) {
		t.Error("If sum and totalNum are same, smaller num should cause begger UCB")
	}
	if calcUcbDefault(10, 20, 100) > calcUcbDefault(10, 20, 200) {
		t.Error("If sum and num are same, bigger totalNum should cause begger UCB")
	}
}
